package topic

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	ExistByID(id uuid.UUID) bool
	GetByID(id uuid.UUID) (*Topic, error)
	GetPage(page int, size int) ([]*Topic, error)
	Create(topic *Topic) error
	Update(topic *Topic) error
}

type InMemoryRepository struct {
	dict map[uuid.UUID]*Topic
	avl  *AVLNode
}

func (im *InMemoryRepository) get(id uuid.UUID) *Topic {
	if im.dict == nil {
		im.dict = make(map[uuid.UUID]*Topic)
	}
	return im.dict[id]
}

func (im *InMemoryRepository) set(id uuid.UUID, topic *Topic) {
	if im.dict == nil {
		im.dict = make(map[uuid.UUID]*Topic)
	}
	im.dict[id] = topic
}

func (im *InMemoryRepository) ExistByID(id uuid.UUID) bool {
	if im.get(id) == nil {
		return false
	}
	return true
}

func (im *InMemoryRepository) GetByID(id uuid.UUID) (*Topic, error) {
	if im.get(id) == nil {
		return nil, fmt.Errorf("TopicNotFound")
	}
	return im.get(id), nil
}

func (im *InMemoryRepository) GetPage(page int, size int) ([]*Topic, error) {
	from := (page-1)*size + 1
	to := from + size - 1
	topics := im.avl.GetRange(from, to)
	return topics, nil
}

func (im *InMemoryRepository) GetAll() ([]*Topic, error) {
	all := im.avl.GetOrderedList()
	return all, nil
}

func (im *InMemoryRepository) Create(topic *Topic) error {
	if im.ExistByID(topic.ID) {
		return fmt.Errorf("TopicAlreadyExist")
	}
	im.set(topic.ID, topic)
	im.avl = im.avl.Insert(topic)
	fmt.Println(im.avl)
	return nil
}

func (im *InMemoryRepository) Update(topic *Topic) error {
	if !im.ExistByID(topic.ID) {
		return fmt.Errorf("TopicNotExist")
	}
	im.avl = im.avl.Delete(topic)
	im.set(topic.ID, topic)
	im.avl = im.avl.Insert(topic)
	return nil
}
