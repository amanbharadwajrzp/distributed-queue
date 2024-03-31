package service

import (
	"mcoding/dao"
	"mcoding/manager"
)

type IWorkerService interface {
	Start() error
	Stop() error
}

type WorkerService struct {
	brokerDao     dao.IBroker
	workerManager manager.IWorkerManager
}

func (w WorkerService) Start() error {
	brokers := w.brokerDao.List()
	for _, broker := range brokers {
		worker := w.workerManager.List(broker)
		func() {

		}
	}
}

func (w WorkerService) Stop() error {
	//TODO implement me
	panic("implement me")
}

func NewWorkerService(workerManager manager.IWorkerManager, brokerDao dao.IBroker) IWorkerService {
	return &WorkerService{
		workerManager: workerManager,
	}
}
