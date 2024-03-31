package main

import (
	"context"
	"fmt"
	"mcoding/boot"
	"mcoding/client"
)

func main() {
	boot.DaoRegistryInstance = boot.RegisterDao()
	boot.RegisterManagers(boot.DaoRegistryInstance)

	broker := client.NewBroker().Create("localhost", 9092)
	broker.AddTopic("topic1", 3)
	broker.AddTopic("topic2", 4)
	broker.Connect()

	producer := client.NewProducer().Broker("localhost", 9092).Topic("topic1")
	err := producer.Publish(client.NewPayload("topic1", "key1", "data1"))
	if err != nil {
		fmt.Println("Error while producing message")
	}

	_ = client.NewWorker().Broker("localhost", 9092).Topic("topic1").Handler(NewWorkerHandler1()).Subscribe()

}

type WorkerHandler1 struct {
}

func (w WorkerHandler1) Handle(ctx context.Context, payload client.IPayload) error {
	fmt.Println("Handling payload from worker 1")
	return nil
}

func (w WorkerHandler1) ZeroPayload() client.IPayload {
	return client.NewPayload("", "", "")
}

func NewWorkerHandler1() client.JobHandler {
	return &WorkerHandler1{}
}
