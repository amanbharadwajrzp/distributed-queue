package client

import (
	"distributed-queue/internal/boot"
	"distributed-queue/internal/dao"
	"distributed-queue/internal/entity"
)

type IBroker interface {
	Create(host string, port int) IBroker
	AddTopic(topic string, numberOfPartition int) IBroker
	Connect() IBroker
}

type Broker struct {
	broker    entity.IBroker
	topics    []entity.Topic
	brokerDao dao.IBroker
	topicDao  dao.ITopic
}

func (b Broker) AddTopic(topicName string, numberOfPartition int) IBroker {
	topic := entity.NewTopic(topicName, numberOfPartition)
	b.topics = append(b.topics, topic)
	return b
}

func (b Broker) Connect() IBroker {
	b.brokerDao.Create(b.broker)
	for _, topic := range b.topics {
		topic.SetBroker(b.broker)
		b.topicDao.SetTopic(topic.GetName(), topic)
	}
	return b
}

func (b Broker) Create(host string, port int) IBroker {
	b.broker = entity.NewBroker(host, port)
	return b
}

func NewBroker() IBroker {
	return &Broker{
		brokerDao: boot.DaoRegistryInstance.GetBrokerDao(),
		topicDao:  boot.DaoRegistryInstance.GetTopicDao(),
	}
}
