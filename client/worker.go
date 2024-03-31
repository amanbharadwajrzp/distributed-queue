package client

import (
	"distributed-queue/internal/boot"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/manager"
)

type Worker struct {
	broker    entity.IBroker
	topic     string
	handler   JobHandler
	workerMgr manager.IWorkerManager
}

func (w Worker) Handler(handler JobHandler) IWorker {
	w.handler = handler
	return w
}

func (w Worker) Broker(host string, port int) IWorker {
	w.broker = entity.NewBroker(host, port)
	return w
}

func (w Worker) Topic(topicName string) IWorker {
	w.topic = topicName
	return w
}

func (w Worker) Subscribe() IWorker {
	workerId := w.workerMgr.RegisterWorker(w.broker)
	w.workerMgr.Start(workerId, w.topic, w.handler)
	return w
}

type IWorker interface {
	Broker(host string, port int) IWorker
	Topic(topicName string) IWorker
	Handler(handler JobHandler) IWorker
	Subscribe() IWorker
}

func NewWorker() IWorker {
	return &Worker{
		workerMgr: boot.ManagerRegistryInstance.GetWorkerManager(),
	}
}
