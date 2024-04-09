package client

import (
	"distributed-queue/internal/boot"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/manager"
)

type IProducer interface {
	Topic(topicName string) IProducer
	Broker(host string, port int) IProducer
	Build() IProducer
	Publish(payload entity.IPayload) error
}

type Producer struct {
	broker       entity.IBroker
	topic        string
	queueManager manager.IQManager
}

func (p Producer) Build() IProducer {
	return p
}

func (p Producer) Topic(topicName string) IProducer {
	p.topic = topicName
	return p
}

func (p Producer) Broker(host string, port int) IProducer {
	p.broker = entity.NewBroker(host, port)
	return p
}

func (p Producer) Publish(payload entity.IPayload) error {
	return p.queueManager.Enqueue(payload)
}

func NewProducer() IProducer {
	return &Producer{
		queueManager: boot.ManagerRegistryInstance.GetQueueManager(),
	}
}
