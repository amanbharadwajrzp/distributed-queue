package manager

import (
	"mcoding/client"
	"mcoding/dao"
	"mcoding/utils"
	"net/rpc"
)

type QueueManager struct {
	queueDao dao.IQueue
	topicDao dao.ITopic
}

func (q QueueManager) Seek(topicName string, partitionNo int, offset int) client.IPayload {
	return q.queueDao.GetPayloads(topicName, partitionNo)[offset]
}

type IQManager interface {
	Enqueue(payload client.IPayload) error
	Dequeue(topicName string) (client.IPayload, error)
	Seek(topicName string, partitionNo int, offset int) client.IPayload
	Purge(topicName string) error
}

func NewQueueManager(queueDao dao.IQueue, topic dao.ITopic) IQManager {
	return &QueueManager{
		queueDao: queueDao,
		topicDao: topic,
	}
}

func (q QueueManager) Enqueue(payload client.IPayload) error {
	if payload == nil || payload.GetTopic() == "" {
		return rpc.ServerError("Payload or Topic cannot be nil")
	}
	topic := q.topicDao.GetTopic(payload.GetTopic())
	if topic == nil {
		return rpc.ServerError("Topic not found")
	}
	q.queueDao.SetPayload(topic.GetName(), utils.GetPartitionNumberFromKey(payload.GetKey(), topic.GetNumberOfPartitions()), payload)
	return nil
}

func (q QueueManager) Dequeue(topicName string) (client.IPayload, error) {
	topic := q.topicDao.GetTopic(topicName)
	queuePayloads := q.queueDao.GetPayloads(topicName, utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions()))
	payload := queuePayloads[len(queuePayloads)-1]
	queuePayloads[len(queuePayloads)-1] = nil
	return payload, nil
}

func (q QueueManager) Purge(topicName string) error {
	topic := q.topicDao.GetTopic(topicName)
	if q.queueDao.GetPartitionPayload(topicName) == nil {
		return rpc.ServerError("Topic not found")
	}
	q.queueDao.GetPartitionPayload(topicName)[utils.GetPartitionNumber(topicName, topic.GetNumberOfPartitions())] = nil
	return nil
}
