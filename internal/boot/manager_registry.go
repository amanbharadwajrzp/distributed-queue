package boot

import (
	"distributed-queue/internal/manager"
)

var (
	ManagerRegistryInstance IManagerRegistry
)

type IManagerRegistry interface {
	GetWorkerManager() manager.IWorkerManager
	GetQueueManager() manager.IQManager
	GetSubscriptionManager() manager.ISubscriptionManager
}

type ManagerRegistry struct {
	workerManager       manager.IWorkerManager
	queueManager        manager.IQManager
	subscriptionManager manager.ISubscriptionManager
}

func (m ManagerRegistry) GetWorkerManager() manager.IWorkerManager {
	return m.workerManager
}

func (m ManagerRegistry) GetQueueManager() manager.IQManager {
	return m.queueManager
}

func (m ManagerRegistry) GetSubscriptionManager() manager.ISubscriptionManager {
	return m.subscriptionManager
}

func RegisterManagers(daoRegistry IDaoRegistry) IManagerRegistry {
	subscriptionManager := manager.NewSubscriptionManager(daoRegistry.GetTopicDao(), daoRegistry.GetSubscriptionDao())
	queueManager := manager.NewQueueManager(daoRegistry.GetQueueDao(), daoRegistry.GetTopicDao())
	workerManager := manager.NewWorkerManager(daoRegistry.GetWorkerDao(), daoRegistry.GetTopicDao(), daoRegistry.GetOffsetDao(), daoRegistry.GetBrokerDao(), subscriptionManager, queueManager)
	return &ManagerRegistry{
		workerManager:       workerManager,
		queueManager:        queueManager,
		subscriptionManager: subscriptionManager,
	}
}
