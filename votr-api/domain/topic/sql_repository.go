package topic

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	uuid "github.com/satori/go.uuid"
)

// Repository layer of Topic domain
type Repository interface {
	ExistByID(id uuid.UUID) bool
	GetByID(id uuid.UUID) (*Topic, error)
	GetPage(keyword string, page int, size int) (*TopicPage, error)
	GetAll() (*TopicPage, error)
	Create(topic *Topic) error
	Update(topic *Topic) error
	DeleteByID(id uuid.UUID) error
}

const (
	querySelectTopic = `
		SELECT
			topic.id,
			topic.title,
			topic.upvote,
			topic.downvote,
			topic.score,
			topic.created,
			topic.updated
		FROM topic
	`
	queryCountTopic = `
		SELECT
			COUNT(topic.id)
		FROM topic
	`
	queryCreateTopic = `
		INSERT INTO topic (
			id,
			title,
			upvote,
			downvote,
			score,
			created,
			updated)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	queryUpdateTopic = `
		UPDATE topic SET
			title = ?,
			upvote = ?,
			downvote = ?,
			score = ?,
			created = ?,
			updated = ?
		WHERE id = ?
	`
	queryDeleteTopic = `
		DELETE FROM topic
		WHERE id = ?
	`
)

type SQLRepository struct {
	DB *sqlx.DB `inject:""`
}

func (r *SQLRepository) ExistByID(id uuid.UUID) (bool, error) {
	var exist bool
	err := r.DB.Get(&exist, queryCountTopic+`
			WHERE topic.id = ?
		`, id)
	if err != nil {
		fmt.Println(err)
	}
	return exist, err
}

func (r *SQLRepository) GetByID(id uuid.UUID) (*Topic, error) {
	var topic Topic
	err := r.DB.Get(&topic, querySelectTopic+`
			WHERE topic.id = ?
		`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("TopicNotFound")
		} else {
			fmt.Println(err)
			return nil, err
		}
	}
	return &topic, err
}

func (r *SQLRepository) GetPage(keyword string, page int, size int) (topicPage *TopicPage, err error) {
	offset := (page - 1) * size
	var topics []*Topic
	var count int
	keywordWildcard := `%` + keyword + `%`
	err = r.DB.Get(&count, queryCountTopic+`
			WHERE title LIKE ?`,
		keywordWildcard)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.DB.Select(&topics, querySelectTopic+`
			WHERE title LIKE ?
			ORDER BY topic.score DESC, created DESC LIMIT ? OFFSET ?`,
		keywordWildcard, size, offset)
	if err != nil {
		fmt.Println(err)
	}
	topicPage = &TopicPage{
		Topics: topics,
		Page:   page,
		Size:   size,
		Total:  count,
	}
	return
}

func (r *SQLRepository) GetAll() (topicPage *TopicPage, err error) {
	var topics []*Topic
	var count int
	err = r.DB.Get(&count, queryCountTopic)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = r.DB.Select(&topics, querySelectTopic+`
			ORDER BY topic.score DESC, created DESC`)
	if err != nil {
		fmt.Println(err)
	}
	topicPage = &TopicPage{
		Topics: topics,
		Page:   1,
		Size:   count,
		Total:  count,
	}
	return
}

func (r *SQLRepository) Create(topic *Topic) (err error) {
	exist, _ := r.ExistByID(topic.ID)
	if exist {
		err = fmt.Errorf("TopicAlreadyExist")
	}
	statementInsert, err := r.DB.Prepare(queryCreateTopic)
	defer statementInsert.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = statementInsert.Exec(
		topic.ID,
		topic.Title,
		topic.Upvote,
		topic.Downvote,
		topic.Score,
		topic.Created,
		topic.Updated,
	)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (r *SQLRepository) Update(topic *Topic) (err error) {
	exist, _ := r.ExistByID(topic.ID)
	if !exist {
		err = fmt.Errorf("TopicNotExist")
	}
	statementUpdate, err := r.DB.Prepare(queryUpdateTopic)
	defer statementUpdate.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = statementUpdate.Exec(
		topic.Title,
		topic.Upvote,
		topic.Downvote,
		topic.Score,
		topic.Created,
		topic.Updated,
		topic.ID,
	)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (r *SQLRepository) DeleteByID(id uuid.UUID) (err error) {
	exist, _ := r.ExistByID(id)
	if !exist {
		err = fmt.Errorf("TopicNotExist")
	}
	statementDel, err := r.DB.Prepare(queryDeleteTopic)
	defer statementDel.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = statementDel.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	return
}
