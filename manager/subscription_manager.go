package manager

import (
	"mcoding/dao"
	"mcoding/utils"
)

type SubscriptionManager struct {
	topic        dao.ITopic
	subscription dao.ISubscription
}

type ISubscriptionManager interface {
	Subscribe(workerID string, topicName string)
	Unsubscribe(workerID string, topicName string)
	Rebalance(topicName string) bool
}

func NewSubscriptionManager(topic dao.ITopic, subscription dao.ISubscription) ISubscriptionManager {
	return &SubscriptionManager{
		topic:        topic,
		subscription: subscription,
	}
}

func (s SubscriptionManager) Subscribe(workerID string, topicName string) {
	topic := s.topic.GetTopic(topicName)
	s.subscription.SetWorkerToTopicsMapping(workerID, topicName)
	s.subscription.SetWorkerToTopicToPartitionMapping(workerID, topicName,
		utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions()))
	s.Rebalance(topicName)
}

func (s SubscriptionManager) Unsubscribe(workerID string, topicName string) {
	s.subscription.DeleteWorkerToTopicsMapping(workerID, topicName)
	s.subscription.DeleteWorkerToTopicToPartitionMapping(workerID, topicName)
}

func (s SubscriptionManager) Rebalance(topicName string) bool {
	topic := s.topic.GetTopic(topicName)
	for workerId, _ := range s.subscription.GetWorkerToTopicsMapping() {
		pNo := utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions())
		s.subscription.SetWorkerToTopicToPartitionMapping(workerId, topicName, pNo)
	}
	return true
}
