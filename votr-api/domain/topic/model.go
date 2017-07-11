package topic

import (
	"fmt"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

type TopicRequest struct {
	Title string `json:"title"`
}

type Topic struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Upvote   int       `json:"upvote"`
	Downvote int       `json:"downvote"`
	Created  time.Time `json:"created"`
	Updated  null.Time `json:"updated"`
}

func NewTopic(name string) *Topic {
	return &Topic{
		ID:      uuid.NewV4(),
		Title:   name,
		Created: time.Now(),
	}
}

func (topic *Topic) Up() {
	topic.Upvote++
	topic.Updated = null.TimeFrom(time.Now())
}

func (topic *Topic) Down() {
	topic.Downvote++
	topic.Updated = null.TimeFrom(time.Now())
}

func (topic *Topic) MoreThan(other *Topic) bool {
	topicScore := topic.Upvote - topic.Downvote
	otherScore := other.Upvote - other.Downvote
	if topicScore == otherScore {
		return topic.Created.After(other.Created)
	}
	return topicScore > otherScore
}

// AVLNode - AVL Tree
// Future: make it more generic
type AVLNode struct {
	topic  *Topic
	left   *AVLNode
	right  *AVLNode
	count  int
	height int
}

func (node *AVLNode) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *AVLNode) Count() int {
	if node == nil {
		return 0
	}
	return node.count
}

// Max return highest score
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewAVLNode(topic *Topic) *AVLNode {
	node := &AVLNode{
		topic:  topic,
		height: 1,
		count:  1,
	}
	return node
}

func (node *AVLNode) rightRotate() *AVLNode {
	nodeY := node
	nodeX := nodeY.left
	nodeT2 := nodeX.right

	nodeX.right = nodeY
	nodeY.left = nodeT2

	nodeY.height = max(nodeY.left.Height(), nodeY.right.Height()) + 1
	nodeX.height = max(nodeX.left.Height(), nodeX.right.Height()) + 1

	nodeY.count = nodeY.left.Count() + nodeY.right.Count() + 1
	nodeX.count = nodeX.left.Count() + nodeX.left.Count() + 1

	return nodeX
}

func (node *AVLNode) leftRotate() *AVLNode {
	nodeX := node
	nodeY := nodeX.right
	nodeT2 := nodeY.left

	nodeY.left = nodeX
	nodeX.right = nodeT2

	nodeX.height = max(nodeX.left.Height(), nodeX.right.Height()) + 1
	nodeY.height = max(nodeY.left.Height(), nodeY.right.Height()) + 1

	nodeY.count = nodeY.left.Count() + nodeY.right.Count() + 1
	nodeX.count = nodeX.left.Count() + nodeX.left.Count() + 1

	return nodeY
}

func (node *AVLNode) getBalance() int {
	if node == nil {
		return 0
	}
	return node.left.Height() - node.right.Height()
}

func (node *AVLNode) Insert(topic *Topic) *AVLNode {
	if node == nil {
		return NewAVLNode(topic)
	}
	if topic.MoreThan(node.topic) {
		node.left = node.left.Insert(topic)
	} else {
		node.right = node.right.Insert(topic)
	}

	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.count = node.left.Count() + node.right.Count() + 1
	balance := node.getBalance()

	// Left left
	if balance > 1 && topic.MoreThan(node.left.topic) {
		return node.rightRotate()
	}

	// Right right
	if balance < -1 && !topic.MoreThan(node.right.topic) {
		return node.leftRotate()
	}

	// Left right
	if balance > 1 && !topic.MoreThan(node.left.topic) {
		node.left = node.left.leftRotate()
		return node.rightRotate()
	}

	// Right left
	if balance < -1 && topic.MoreThan(node.right.topic) {
		node.right = node.right.rightRotate()
		return node.leftRotate()
	}
	return node
}

func (node *AVLNode) maxValueNode() *AVLNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (node *AVLNode) Delete(topic *Topic) *AVLNode {
	// 1. standard bst delete
	if node == nil {
		return node
	}

	if uuid.Equal(topic.ID, node.topic.ID) {
		if node.left == nil || node.right == nil {
			temp := node.left
			if temp == nil {
				temp = node.right
			}

			// 0 child
			if temp == nil {
				node = nil
			} else { // 1 child
				node = temp
			}
		} else {
			temp := node.maxValueNode()
			node.topic = temp.topic
			node.right.Delete(temp.topic)
		}
	} else if topic.MoreThan(node.topic) {
		node.left = node.left.Delete(topic)
	} else {
		node.right = node.right.Delete(topic)
	}

	if node == nil {
		return node
	}

	// 2. Update height
	node.height = max(node.left.Height(), node.right.Height()) + 1
	node.count = node.left.Count() + node.right.Count() + 1

	balance := node.getBalance()

	// Left left
	if balance > 1 && node.left.getBalance() >= 0 {
		return node.rightRotate()
	}

	// Left right
	if balance > 1 && node.left.getBalance() < 0 {
		node.left = node.left.leftRotate()
		return node.rightRotate()
	}

	// Right right
	if balance < -1 && node.right.getBalance() <= 0 {
		return node.leftRotate()
	}

	// Right left
	if balance < 1 && node.right.getBalance() > 0 {
		node.right = node.right.rightRotate()
		return node.leftRotate()
	}
	return node
}

func (node *AVLNode) GetOrderedList() []*Topic {
	if node != nil {
		left := node.left.GetOrderedList()
		right := node.right.GetOrderedList()
		tmp := append(left, node.topic)
		return append(tmp, right...)
	}
	return []*Topic{}
}

func (node *AVLNode) GetRange(from int, to int) []*Topic {
	fmt.Println("GetRange", from, to)
	if to < from {
		return []*Topic{}
	}
	if node != nil {
		var left, right []*Topic
		if node.left.Count() >= from { // start from left
			// fmt.Println("A", node.left.Count(), node.right.Count())
			left = node.left.GetRange(from, to)
		}

		if node.left.Count() >= from-1 && to-node.left.Count()-1 >= 0 {
			// fmt.Println("B", node.left.Count(), node.right.Count())
			left = append(left, node.topic)
		}

		if to-node.left.Count()-1 > 0 {
			// fmt.Println("C", node.left.Count(), node.right.Count())
			fromLocal := from - node.left.Count() - 1
			toLocal := to - node.left.Count() - 1
			right = node.right.GetRange(fromLocal, toLocal)
		}
		return append(left, right...)
	}
	return []*Topic{}
}
