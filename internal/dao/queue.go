package dao

import (
	"distributed-queue/internal/entity"
)

type Queue struct {
	topicToPartitionToPayloads map[string]map[int][]entity.IPayload
}

func (q Queue) Initialize(topic string, partitionNumber int) {
	if q.topicToPartitionToPayloads[topic] == nil {
		q.topicToPartitionToPayloads[topic] = make(map[int][]entity.IPayload)
	}
	if q.topicToPartitionToPayloads[topic][partitionNumber] == nil {
		q.topicToPartitionToPayloads[topic][partitionNumber] = make([]entity.IPayload, 0)
	}
}

func (q Queue) GetPartitionPayload(topic string) map[int][]entity.IPayload {
	return q.topicToPartitionToPayloads[topic]
}

func (q Queue) GetPayloads(topic string, partitionNumber int) []entity.IPayload {
	if q.topicToPartitionToPayloads[topic] == nil {
		return nil
	}
	return q.topicToPartitionToPayloads[topic][partitionNumber]
}

func (q Queue) SetPayload(topic string, partitionNumber int, payload entity.IPayload) {
	if q.topicToPartitionToPayloads[topic] == nil {
		q.topicToPartitionToPayloads[topic] = make(map[int][]entity.IPayload)
	}
	if q.topicToPartitionToPayloads[topic][partitionNumber] == nil {
		q.topicToPartitionToPayloads[topic][partitionNumber] = make([]entity.IPayload, 0)
	}
	q.topicToPartitionToPayloads[topic][partitionNumber] = append(q.topicToPartitionToPayloads[topic][partitionNumber], payload)
}

type IQueue interface {
	GetPartitionPayload(topic string) map[int][]entity.IPayload
	GetPayloads(topic string, partitionNumber int) []entity.IPayload
	SetPayload(topic string, partitionNumber int, payload entity.IPayload)
	Initialize(topic string, partitions int)
}

func NewQueue() IQueue {
	return &Queue{
		topicToPartitionToPayloads: make(map[string]map[int][]entity.IPayload),
	}
}
