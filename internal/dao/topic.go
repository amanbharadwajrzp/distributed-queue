package dao

import "distributed-queue/internal/entity"

var (
	topicNameToTopicMapping map[string]entity.Topic
)

type Topic struct {
}

func (t Topic) GetTopic(topicName string) entity.Topic {
	return topicNameToTopicMapping[topicName]
}

func (t Topic) SetTopic(topicName string, topic entity.Topic) {
	if topicNameToTopicMapping == nil {
		topicNameToTopicMapping = make(map[string]entity.Topic)
	}
	topicNameToTopicMapping[topicName] = topic
}

type ITopic interface {
	GetTopic(topicName string) entity.Topic
	SetTopic(topicName string, topic entity.Topic)
}

func NewTopic() ITopic {
	return &Topic{}
}
