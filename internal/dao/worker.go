package dao

import (
	"distributed-queue/internal/entity"
	"distributed-queue/internal/utils"
)

type Worker struct {
	workerIdToWorkerMap map[string]entity.IWorker
	brokerToWorkerMap   map[string][]entity.IWorker
}

type IWorker interface {
	Get(workerId string) entity.IWorker
	Add(broker entity.IBroker, workerId string) entity.IWorker
	Delete(broker entity.IBroker, workerId string) bool
	List(broker entity.IBroker) []entity.IWorker
}

func NewWorker() IWorker {
	return &Worker{
		workerIdToWorkerMap: make(map[string]entity.IWorker),
		brokerToWorkerMap:   make(map[string][]entity.IWorker),
	}
}

func (w Worker) Get(workerId string) entity.IWorker {
	return w.workerIdToWorkerMap[workerId]
}

func (w Worker) Add(broker entity.IBroker, workerId string) entity.IWorker {
	brokerId := entity.GetBrokerId(broker.GetHost(), broker.GetPort())
	worker := entity.NewWorker(workerId, broker)
	w.workerIdToWorkerMap[worker.GetWorkerId()] = worker
	if w.brokerToWorkerMap[brokerId] == nil {
		w.brokerToWorkerMap[brokerId] = make([]entity.IWorker, 0)
	}
	w.brokerToWorkerMap[brokerId] = append(w.brokerToWorkerMap[brokerId], worker)
	return worker
}

func (w Worker) List(broker entity.IBroker) []entity.IWorker {
	return w.brokerToWorkerMap[entity.GetBrokerId(broker.GetHost(), broker.GetPort())]
}

func (w Worker) Delete(broker entity.IBroker, workerId string) bool {
	brokerId := entity.GetBrokerId(broker.GetHost(), broker.GetPort())
	w.workerIdToWorkerMap[workerId] = nil
	workerInterface := make([]interface{}, 0)
	for _, worker := range w.brokerToWorkerMap[brokerId] {
		workerInterface = append(workerInterface, worker)
	}
	removedWorkerSlice := utils.RemoveFromSliceInterface(workerInterface, workerId)
	w.brokerToWorkerMap[brokerId] = make([]entity.IWorker, 0)
	for _, worker := range removedWorkerSlice {
		w.brokerToWorkerMap[brokerId] = append(w.brokerToWorkerMap[brokerId], worker.(entity.IWorker))
	}
	return true
}
