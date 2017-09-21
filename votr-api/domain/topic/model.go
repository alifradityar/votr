package topic

import (
	"fmt"
	"strings"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

type TopicRequest struct {
	Title string `json:"title"`
}

type Topic struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Upvote   int       `json:"upvote" db:"upvote"`
	Downvote int       `json:"downvote" db:"downvote"`
	Score    int       `json:"score" db:"score"`
	Created  time.Time `json:"created" db:"created"`
	Updated  null.Time `json:"updated" db:"updated"`
}

func NewTopic(name string) (*Topic, error) {
	if len(name) < 3 || len(name) > 255 {
		return nil, fmt.Errorf("InvalidTopicTitleLength")
	}
	return &Topic{
		ID:      uuid.NewV4(),
		Title:   strings.TrimSpace(name),
		Created: time.Now(),
	}, nil
}

func (topic *Topic) SetTitle(title string) {
	topic.Title = title
	topic.Updated = null.TimeFrom(time.Now())
}

func (topic *Topic) Up() {
	topic.Upvote++
	topic.Score++
	topic.Updated = null.TimeFrom(time.Now())
}

func (topic *Topic) Down() {
	topic.Downvote++
	topic.Score--
	topic.Updated = null.TimeFrom(time.Now())
}

func (topic *Topic) MoreThan(other *Topic) bool {
	return topic.Score > other.Score
}

type TopicPage struct {
	Topics []*Topic `json:"topics"`
	Page   int      `json:"page"`
	Size   int      `json:"size"`
	Total  int      `json:"total"`
}

func NewTopicPage(topics []*Topic, page int, total int) *TopicPage {
	return &TopicPage{
		Topics: topics,
		Page:   page,
		Size:   len(topics),
		Total:  total,
	}
}
