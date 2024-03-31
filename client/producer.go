package client

import (
	"distributed-queue/internal/boot"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/manager"
)

type IProducer interface {
	Topic(topicName string) IProducer
	Broker(host string, port int) IProducer
	Publish(payload IPayload) error
}

type Producer struct {
	broker       entity.IBroker
	topic        string
	queueManager manager.IQManager
}

func (p Producer) Topic(topicName string) IProducer {
	p.topic = topicName
	return p
}

func (p Producer) Broker(host string, port int) IProducer {
	p.broker = entity.NewBroker(host, port)
	return p
}

func (p Producer) Publish(payload IPayload) error {
	payload.SetTopic(p.topic)
	return p.queueManager.Enqueue(payload)
}

func NewProducer() IProducer {
	return &Producer{
		queueManager: boot.ManagerRegistryInstance.GetQueueManager(),
	}
}
