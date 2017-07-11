package topic

import uuid "github.com/satori/go.uuid"

type Service struct {
	Repository *InMemoryRepository `inject:""`
}

func (s *Service) CreateTopic(topicName string) (*Topic, error) {
	topic := NewTopic(topicName)
	err := s.Repository.Create(topic)
	return topic, err
}

func (s *Service) UpvoteTopic(topicID uuid.UUID) (*Topic, error) {
	topic, err := s.Repository.GetByID(topicID)
	if err != nil {
		return topic, err
	}
	topic.Up()
	err = s.Repository.Update(topic)
	return topic, err
}

func (s *Service) DownvoteTopic(topicID uuid.UUID) (*Topic, error) {
	topic, err := s.Repository.GetByID(topicID)
	if err != nil {
		return topic, err
	}
	topic.Down()
	err = s.Repository.Update(topic)
	return topic, err
}

func (s *Service) GetTopic(topicID uuid.UUID) (*Topic, error) {
	return s.Repository.GetByID(topicID)
}

func (s *Service) GetTopicPage(page int, size int) ([]*Topic, error) {
	return s.Repository.GetPage(page, size)
}

func (s *Service) GetAll() ([]*Topic, error) {
	return s.Repository.GetAll()
}
