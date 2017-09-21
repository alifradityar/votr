package topic

import uuid "github.com/satori/go.uuid"

type Service struct {
	Repository *CacheRepository `inject:""`
}

func (s *Service) CreateTopic(topicName string) (*Topic, error) {
	topic, err := NewTopic(topicName)
	if err != nil {
		return topic, err
	}
	err = s.Repository.Create(topic)
	return topic, err
}

func (s *Service) UpdateTopic(topicID uuid.UUID, topicName string) (*Topic, error) {
	topic, err := s.Repository.GetByID(topicID)
	if err != nil {
		return topic, err
	}
	topic.SetTitle(topicName)
	err = s.Repository.Update(topic)
	return topic, err
}

func (s *Service) DeleteTopic(topicID uuid.UUID) error {
	return s.Repository.DeleteByID(topicID)
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

func (s *Service) GetTopicPage(keyword string, page int, size int) (*TopicPage, error) {
	return s.Repository.GetPage(keyword, page, size)
}

func (s *Service) GetAll() (*TopicPage, error) {
	return s.Repository.GetAll()
}
