package boot

import "distributed-queue/internal/dao"

var (
	DaoRegistryInstance IDaoRegistry
)

type IDaoRegistry interface {
	GetWorkerDao() dao.IWorker
	GetBrokerDao() dao.IBroker
	GetTopicDao() dao.ITopic
	GetOffsetDao() dao.IOffset
	GetQueueDao() dao.IQueue
	GetSubscriptionDao() dao.ISubscription
}

type DaoRegistry struct {
	workerDao       dao.IWorker
	brokerDao       dao.IBroker
	topicDao        dao.ITopic
	offsetDao       dao.IOffset
	queueDao        dao.IQueue
	subscriptionDao dao.ISubscription
}

func (d DaoRegistry) GetWorkerDao() dao.IWorker {
	return d.workerDao
}

func (d DaoRegistry) GetBrokerDao() dao.IBroker {
	return d.brokerDao
}

func (d DaoRegistry) GetTopicDao() dao.ITopic {
	return d.topicDao
}

func (d DaoRegistry) GetOffsetDao() dao.IOffset {
	return d.offsetDao
}

func (d DaoRegistry) GetQueueDao() dao.IQueue {
	return d.queueDao
}

func (d DaoRegistry) GetSubscriptionDao() dao.ISubscription {
	return d.subscriptionDao
}

func RegisterDao() IDaoRegistry {
	return &DaoRegistry{
		workerDao:       dao.NewWorker(),
		brokerDao:       dao.NewBroker(),
		topicDao:        dao.NewTopic(),
		offsetDao:       dao.NewOffset(),
		queueDao:        dao.NewQueue(),
		subscriptionDao: dao.NewSubscription(),
	}
}
