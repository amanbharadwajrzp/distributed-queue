package dao

import (
	"distributed-queue/client"
)

type Queue struct {
	topicToPartitionToPayloads map[string]map[int][]client.IPayload
}

func (q Queue) GetPartitionPayload(topic string) map[int][]client.IPayload {
	return q.topicToPartitionToPayloads[topic]
}

func (q Queue) GetPayloads(topic string, partitionNumber int) []client.IPayload {
	if q.topicToPartitionToPayloads[topic] == nil {
		return nil
	}
	return q.topicToPartitionToPayloads[topic][partitionNumber]
}

func (q Queue) SetPayload(topic string, partitionNumber int, payload client.IPayload) {
	if q.topicToPartitionToPayloads[topic] == nil {
		q.topicToPartitionToPayloads[topic] = make(map[int][]client.IPayload)
	}
	if q.topicToPartitionToPayloads[topic][partitionNumber] == nil {
		q.topicToPartitionToPayloads[topic][partitionNumber] = make([]client.IPayload, 0)
	}
	q.topicToPartitionToPayloads[topic][partitionNumber] = append(q.topicToPartitionToPayloads[topic][partitionNumber], payload)
}

type IQueue interface {
	GetPartitionPayload(topic string) map[int][]client.IPayload
	GetPayloads(topic string, partitionNumber int) []client.IPayload
	SetPayload(topic string, partitionNumber int, payload client.IPayload)
}

func NewQueue() IQueue {
	return &Queue{
		topicToPartitionToPayloads: make(map[string]map[int][]client.IPayload),
	}
}
