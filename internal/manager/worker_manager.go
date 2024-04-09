package manager

import (
	"context"
	"distributed-queue/internal/dao"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/utils"
	"fmt"
	"time"
)

type WorkerManager struct {
	brokerDao       dao.IBroker
	workerDao       dao.IWorker
	topicsDao       dao.ITopic
	offsetDao       dao.IOffset
	subscriptionMgr ISubscriptionManager
	queueMgr        IQManager
	poller          *Poller
}

type Poller struct {
	ticker time.Ticker
	done   chan struct{}
}

func (w WorkerManager) GetWorker(workerId string) entity.IWorker {
	return w.workerDao.Get(workerId)
}

func (w WorkerManager) List(broker entity.IBroker) []entity.IWorker {
	return w.workerDao.List(broker)
}

func (w WorkerManager) RegisterWorker(broker entity.IBroker, workerId string) {
	_ = w.workerDao.Add(broker, workerId)
}

func (w WorkerManager) UnregisterWorker(workerId string, brokerId string) bool {
	broker := w.brokerDao.Get(brokerId)
	w.workerDao.Delete(broker, workerId)
	return true
}

func (w WorkerManager) Start(workerId string, topicName string, handler entity.JobHandler) bool {
	w.subscriptionMgr.Subscribe(workerId, topicName)
	topic := w.topicsDao.GetTopic(topicName)
	pNo := utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions())
	w.offsetDao.Create(workerId, topicName, pNo)
	go func() {
		for {
			select {
			case <-w.poller.done:
				return
			case <-w.poller.ticker.C:
				w.pickLatest(workerId, topicName, pNo, handler)
			}
		}
	}()
	return true
}

func (w WorkerManager) pickLatest(workerId string, topicName string, pNo int, handler entity.JobHandler) {
	offset := w.offsetDao.GetLatest(workerId, topicName, pNo)
	payload := w.queueMgr.Seek(topicName, pNo, offset.Offset)
	if payload != nil {
		fmt.Println("Worker", workerId, "picked new payload")
		err := handler.Handle(context.Background(), payload)
		if err != nil {
			w.offsetDao.NAck(workerId, topicName, pNo)
		} else {
			w.offsetDao.Ack(workerId, topicName, pNo)
		}
	} else {
		fmt.Println("Worker", workerId, "is waiting for new payload")
	}
}

func (w WorkerManager) Stop() bool {
	w.poller.ticker.Stop()
	w.poller.done <- struct{}{}
	return true
}

type IWorkerManager interface {
	RegisterWorker(broker entity.IBroker, workerId string)
	UnregisterWorker(workerId string, brokerId string) bool
	Start(workerId string, topicName string, handler entity.JobHandler) bool
	Stop() bool
	GetWorker(workerId string) entity.IWorker
	List(broker entity.IBroker) []entity.IWorker
}

func NewWorkerManager(workerDao dao.IWorker, topicsDao dao.ITopic, offsetDao dao.IOffset, brokerDao dao.IBroker, subscriptionMgr ISubscriptionManager, queueMgr IQManager) IWorkerManager {
	poller := &Poller{
		ticker: *time.NewTicker(1 * time.Second),
		done:   make(chan struct{}),
	}
	return &WorkerManager{
		brokerDao:       brokerDao,
		workerDao:       workerDao,
		topicsDao:       topicsDao,
		offsetDao:       offsetDao,
		subscriptionMgr: subscriptionMgr,
		queueMgr:        queueMgr,
		poller:          poller,
	}
}
