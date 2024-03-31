package main

import (
	"distributed-queue/client"
	"distributed-queue/e2e/worker"
	"distributed-queue/internal/boot"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// Register the handlers
	for topic, handler := range worker.GetTopicToHandlersMapping() {
		_ = client.NewWorker().Broker(boot.Config.Broker.Hostname, boot.Config.Broker.Port).Topic(topic).Handler(handler).Subscribe()
	}

	http.HandleFunc("/worker/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Worker running fine! on %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":9040", nil))
}
