package main

import (
	"distributed-queue/client"
	"distributed-queue/cmd/producer"
	"distributed-queue/cmd/worker"
	"distributed-queue/internal/boot"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	boot.Initialize()
	router := gin.Default()
	RegisterBroker(router)
	RegisterWorker()
	RegisterProducer(router)
	err := router.Run(fmt.Sprintf(":%d", boot.GetAppConfig().Port))
	if err != nil {
		log.Fatal("Unable to start server. Error: ", err)
	}
	shutdownWorkers()
}

func shutdownWorkers() {
	c := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM.
	// SIGKILL, SIGQUIT will not be caught.
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	// Block until signal is received.
	<-c
	fmt.Println("Stopping workers")
	for _, w := range worker.GetRunningWorkers() {
		w.Stop()
	}
}

func RegisterProducer(router *gin.Engine) {
	router.GET("/producer/check", func(c *gin.Context) {
		fmt.Println("Producer running fine!")
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.POST("/producer/publish/:topic", producer.ProducerHandlerFunc)
}

func RegisterWorker() {
	// Register the handlers
	for topic, handler := range worker.GetTopicToHandlersMapping() {
		w := client.NewWorker().Broker(boot.GetAppConfig().Hostname, boot.GetAppConfig().Port).Topic(topic).Handler(handler).Subscribe()
		worker.AddRunningWorker(w)
	}
}

func RegisterBroker(router *gin.Engine) {
	topics := boot.Config.App.Topics
	broker := client.NewBroker().Create(boot.GetAppConfig().Hostname, boot.GetAppConfig().Port)
	for _, topic := range topics {
		broker.AddTopic(topic.Name, topic.Partitions)
	}
	broker.Connect()
	router.GET("/broker/check", func(c *gin.Context) {
		fmt.Println("Broker running fine!")
		c.JSON(200, gin.H{"status": "ok"})
	})
}
