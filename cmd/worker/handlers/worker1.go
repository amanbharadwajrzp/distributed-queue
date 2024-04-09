package handlers

import (
	"context"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/utils"
	"fmt"
)

type Handler1 struct {
}

func (w Handler1) Name() string {
	return "Handler1"
}

func (w Handler1) Handle(ctx context.Context, payload entity.IPayload) error {
	fmt.Println("Handling payload from worker 1")
	return nil
}

func (w Handler1) ZeroPayload() entity.IPayload {
	return entity.NewPayload("", "", nil, utils.String)
}

func NewHandler1() entity.JobHandler {
	return &Handler1{}
}
