package dao

import "mcoding/entity"

type Topic struct {
	topicNameToTopicMapping map[string]entity.Topic
}

func (t Topic) GetTopic(topicName string) entity.Topic {
	return t.topicNameToTopicMapping[topicName]
}

func (t Topic) SetTopic(topicName string, topic entity.Topic) {
	t.topicNameToTopicMapping[topicName] = topic
}

type ITopic interface {
	GetTopic(topicName string) entity.Topic
	SetTopic(topicName string, topic entity.Topic)
}

func NewTopic() ITopic {
	return &Topic{
		topicNameToTopicMapping: make(map[string]entity.Topic),
	}
}
