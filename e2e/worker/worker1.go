package worker

import (
	"context"
	"distributed-queue/client"
	"fmt"
)

type Handler1 struct {
}

func (w Handler1) Handle(ctx context.Context, payload client.IPayload) error {
	fmt.Println("Handling payload from worker 1")
	return nil
}

func (w Handler1) ZeroPayload() client.IPayload {
	return client.NewPayload("", "", "")
}

func NewHandler1() client.JobHandler {
	return &Handler1{}
}
