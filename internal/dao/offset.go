package dao

import "distributed-queue/internal/entity"

type Offset struct {
	workerToTopicToPartitionToOffset map[string]map[string]map[int]*entity.Offset
}

func (o Offset) Create(workerId string, topicName string, partition int) bool {
	if o.workerToTopicToPartitionToOffset[workerId] == nil {
		o.workerToTopicToPartitionToOffset[workerId] = make(map[string]map[int]*entity.Offset)
	}
	if o.workerToTopicToPartitionToOffset[workerId][topicName] == nil {
		o.workerToTopicToPartitionToOffset[workerId][topicName] = make(map[int]*entity.Offset)
	}
	o.workerToTopicToPartitionToOffset[workerId][topicName][partition] = &entity.Offset{Offset: -1, Status: entity.NOT_PICKED}
	return true
}

func (o Offset) Delete(workerId string, topicName string, partition int) bool {
	if o.workerToTopicToPartitionToOffset[workerId] == nil ||
		o.workerToTopicToPartitionToOffset[workerId][topicName] == nil {
		return false
	}
	o.workerToTopicToPartitionToOffset[workerId][topicName][partition] = nil
	return true
}

func (o Offset) GetLatest(workerId string, topicName string, partition int) *entity.Offset {
	if o.workerToTopicToPartitionToOffset[workerId] == nil ||
		o.workerToTopicToPartitionToOffset[workerId][topicName] == nil {
		return nil
	}
	return o.workerToTopicToPartitionToOffset[workerId][topicName][partition]
}

func (o Offset) Ack(workerId string, topicName string, partition int) bool {
	if o.workerToTopicToPartitionToOffset[workerId] == nil ||
		o.workerToTopicToPartitionToOffset[workerId][topicName] == nil {
		return false
	}
	o.workerToTopicToPartitionToOffset[workerId][topicName][partition].Offset++
	return true
}

func (o Offset) NAck(workerId string, topicName string, partition int) bool {
	if o.workerToTopicToPartitionToOffset[workerId] == nil ||
		o.workerToTopicToPartitionToOffset[workerId][topicName] == nil {
		return false
	}
	o.workerToTopicToPartitionToOffset[workerId][topicName][partition].Status = entity.NACK_STATUS
	return true
}

type IOffset interface {
	Create(workerId string, topicName string, partition int) bool
	Delete(workerId string, topicName string, partition int) bool
	GetLatest(workerId string, topicName string, partition int) *entity.Offset
	Ack(workerId string, topicName string, partition int) bool
	NAck(workerId string, topicName string, partition int) bool
}

func NewOffset() IOffset {
	return &Offset{
		workerToTopicToPartitionToOffset: make(map[string]map[string]map[int]*entity.Offset),
	}
}
