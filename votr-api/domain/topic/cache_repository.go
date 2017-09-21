package topic

import (
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

// CacheRepository is an implementation of repository layer of Topic Domain
type CacheRepository struct {
	RedisClient   *redis.Client  `inject:""`
	SQLRepository *SQLRepository `inject:""`
}

func (c *CacheRepository) getSingleKey(id uuid.UUID) string {
	return fmt.Sprintf("topic:%s", id.String())
}

func (c *CacheRepository) getListKey() string {
	return fmt.Sprintf("topicList")
}

func (c *CacheRepository) syncSingle(topic *Topic) error {
	topicJSON, _ := json.Marshal(*topic)
	topicString := string(topicJSON)
	err := c.RedisClient.ZAdd(c.getListKey(), redis.Z{
		Score:  float64(topic.Score),
		Member: topic.ID.String(),
	}).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = c.RedisClient.Set(c.getSingleKey(topic.ID), topicString, 0).Err()
	return err
}

func (c *CacheRepository) ExistByID(id uuid.UUID) (bool, error) {
	existInt, err := c.RedisClient.Exists(c.getSingleKey(id)).Result()
	exist := existInt > 0
	if err == redis.Nil {
		err = nil
	}
	if !exist {
		exist, err = c.SQLRepository.ExistByID(id)
		if exist {
			// Sync sql to cache
			go func() {
				topic, err := c.SQLRepository.GetByID(id)
				if err != nil {
					return
				}
				c.syncSingle(topic)
			}()
		}
		return exist, err
	}
	return exist, err
}

func (c *CacheRepository) GetByID(id uuid.UUID) (*Topic, error) {
	var topic Topic
	redisResult, err := c.RedisClient.Get(c.getSingleKey(id)).Result()
	if err == redis.Nil {
		topic, err := c.SQLRepository.GetByID(id)
		if topic != nil {
			go c.syncSingle(topic)
		}
		return nil, err
	}
	err = json.Unmarshal([]byte(redisResult), &topic)
	return &topic, err
}

func (c *CacheRepository) GetPage(keyword string, page int, size int) (*TopicPage, error) {
	from := int64((page - 1) * size)
	to := from + int64(size)
	topicIDs, err := c.RedisClient.ZRevRange(c.getListKey(), from, to).Result()
	if err != nil {
		return nil, err
	}
	count, err := c.RedisClient.ZCount(c.getListKey(), "-inf", "+inf").Result()
	var topics []*Topic
	for _, topicID := range topicIDs {
		topic, _ := c.GetByID(uuid.FromStringOrNil(topicID))
		topics = append(topics, topic)
	}
	return &TopicPage{
		Topics: topics,
		Page:   page,
		Size:   size,
		Total:  int(count),
	}, nil
}

func (c *CacheRepository) GetAll() (*TopicPage, error) {
	topicIDs, err := c.RedisClient.ZRevRange(c.getListKey(), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	count, err := c.RedisClient.ZCount(c.getListKey(), "-inf", "+inf").Result()
	var topics []*Topic
	for _, topicID := range topicIDs {
		topic, _ := c.GetByID(uuid.FromStringOrNil(topicID))
		topics = append(topics, topic)
	}
	return &TopicPage{
		Topics: topics,
		Page:   1,
		Size:   int(count),
		Total:  int(count),
	}, nil
}

func (c *CacheRepository) Create(topic *Topic) error {
	err := c.SQLRepository.Create(topic)
	go c.syncSingle(topic)
	return err
}

func (c *CacheRepository) Update(topic *Topic) error {
	err := c.SQLRepository.Update(topic)
	go c.syncSingle(topic)
	return err
}

func (c *CacheRepository) DeleteByID(id uuid.UUID) error {
	err := c.SQLRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	err = c.RedisClient.Del(c.getSingleKey(id)).Err()
	if err != nil {
		return err
	}
	err = c.RedisClient.ZRem(c.getListKey(), id.String()).Err()
	return err
}
