package worker

import "distributed-queue/client"

var (
	topicsToHandlersMapping = map[string]client.JobHandler{
		"topic1": NewHandler1(),
	}
)

func GetTopicToHandlersMapping() map[string]client.JobHandler {
	return topicsToHandlersMapping
}
