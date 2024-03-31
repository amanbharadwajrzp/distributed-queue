package main

import (
	"distributed-queue/client"
	"distributed-queue/internal/boot"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	boot.Initialize()
	brokerConfig := boot.Config.Broker
	topics := brokerConfig.Topics
	broker := client.NewBroker().Create(brokerConfig.Hostname, brokerConfig.Port)
	for _, topic := range topics {
		broker.AddTopic(topic.Name, topic.Partitions)
	}
	broker.Connect()
	http.HandleFunc("/broker/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Broker running fine! on %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", brokerConfig.Port), nil))
}
