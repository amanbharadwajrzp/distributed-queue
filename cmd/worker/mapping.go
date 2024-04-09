package worker

import (
	"distributed-queue/client"
	"distributed-queue/cmd/worker/handlers"
	"distributed-queue/internal/entity"
)

var (
	topicsToHandlersMapping = map[string]entity.JobHandler{
		"topic1": handlers.NewHandler1(),
	}
	workers = make([]client.IWorker, 0)
)

func GetTopicToHandlersMapping() map[string]entity.JobHandler {
	return topicsToHandlersMapping
}
func GetRunningWorkers() []client.IWorker {
	return workers
}

func AddRunningWorker(worker client.IWorker) {
	workers = append(workers, worker)
}
