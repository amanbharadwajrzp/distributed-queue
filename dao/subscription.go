package dao

import "mcoding/utils"

type Subscription struct {
	workerToTopicsMapping           map[string][]string
	workerToTopicToPartitionMapping map[string]map[string]int
}

func (s Subscription) DeleteWorkerToTopicsMapping(workerId string, topic string) bool {
	s.workerToTopicsMapping[workerId] = utils.RemoveFromSlice(s.workerToTopicsMapping[workerId], topic)
	return true
}

func (s Subscription) DeleteWorkerToTopicToPartitionMapping(workerId string, topicName string) bool {
	s.workerToTopicToPartitionMapping[workerId] = nil
	return true
}

func (s Subscription) SetWorkerToTopicsMapping(workerId string, topic string) bool {
	if s.workerToTopicsMapping[workerId] == nil || len(s.workerToTopicsMapping[workerId]) == 0 {
		s.workerToTopicsMapping[workerId] = make([]string, 0)
	}
	s.workerToTopicsMapping[workerId] = append(s.workerToTopicsMapping[workerId], topic)
	return true
}

func (s Subscription) SetWorkerToTopicToPartitionMapping(workerId string, topicName string, partition int) bool {
	if s.workerToTopicToPartitionMapping[workerId] == nil {
		s.workerToTopicToPartitionMapping[workerId] = make(map[string]int)
	}
	s.workerToTopicToPartitionMapping[workerId][topicName] = partition
	return true
}

func (s Subscription) GetWorkerToTopicsMapping() map[string][]string {
	return s.workerToTopicsMapping
}

func (s Subscription) GetWorkerToTopicToPartitionMapping() map[string]map[string]int {
	return s.workerToTopicToPartitionMapping
}

type ISubscription interface {
	GetWorkerToTopicsMapping() map[string][]string
	GetWorkerToTopicToPartitionMapping() map[string]map[string]int
	SetWorkerToTopicsMapping(workerId string, topic string) bool
	SetWorkerToTopicToPartitionMapping(workerId string, topicName string, partition int) bool
	DeleteWorkerToTopicsMapping(workerId string, topic string) bool
	DeleteWorkerToTopicToPartitionMapping(workerId string, topicName string) bool
}

func NewSubscription() ISubscription {
	return &Subscription{
		workerToTopicsMapping: make(map[string][]string),
	}
}
