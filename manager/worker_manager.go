package manager

import (
	"context"
	"mcoding/dao"
	"mcoding/entity"
	"mcoding/utils"
)

type WorkerManager struct {
	brokerDao       dao.IBroker
	workerDao       dao.IWorker
	topicsDao       dao.ITopic
	offsetDao       dao.IOffset
	subscriptionMgr ISubscriptionManager
	queueMgr        IQManager
}

func (w WorkerManager) GetWorker(workerId string) entity.IWorker {
	return w.workerDao.Get(workerId)
}

func (w WorkerManager) List(broker entity.IBroker) []entity.IWorker {
	return w.workerDao.List(broker)
}

func (w WorkerManager) RegisterWorker(broker entity.IBroker) string {
	worker := w.workerDao.Add(broker)
	return worker.GetWorkerId()
}

func (w WorkerManager) UnregisterWorker(workerId string, brokerId string) bool {
	broker := w.brokerDao.Get(brokerId)
	w.workerDao.Delete(*broker, workerId)
	return true
}

func (w WorkerManager) Start(workerId string, topicName string, handler entity.JobHandler) bool {
	w.subscriptionMgr.Subscribe(workerId, topicName)
	topic := w.topicsDao.GetTopic(topicName)
	pNo := utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions())
	w.offsetDao.Create(workerId, topicName, pNo)
	go func() {
		for {
			offset := w.offsetDao.GetLatest(workerId, topicName, pNo)
			payload := w.queueMgr.Seek(topicName, pNo, offset.Offset)
			if payload != nil {
				err := handler.Handle(context.Background(), payload)
				if err != nil {
					w.offsetDao.NAck(workerId, topicName, pNo)
				} else {
					w.offsetDao.Ack(workerId, topicName, pNo)
				}
			}
		}
	}()
	return true
}

type IWorkerManager interface {
	RegisterWorker(broker entity.IBroker) string
	UnregisterWorker(workerId string, brokerId string) bool
	Start(workerId string, topicName string, handler entity.JobHandler) bool
	GetWorker(workerId string) entity.IWorker
	List(broker entity.IBroker) []entity.IWorker
}

func NewWorkerManager(workerDao dao.IWorker, topicsDao dao.ITopic, offsetDao dao.IOffset, brokerDao dao.IBroker, subscriptionMgr ISubscriptionManager, queueMgr IQManager) IWorkerManager {
	return &WorkerManager{
		brokerDao:       brokerDao,
		workerDao:       workerDao,
		topicsDao:       topicsDao,
		offsetDao:       offsetDao,
		subscriptionMgr: subscriptionMgr,
		queueMgr:        queueMgr,
	}
}
