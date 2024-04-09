# Distrubuted In-memory Queue.
A simple in-memory queue that supports partitioning and offset management. Multiple consumer workers can subscribe to a single topic.

#### Local Setup Setup broker and worker servers.

- Run `make go-run-api` to run api server.
	- It runs `make go-build-api` that builds api binary.


#### Register worker handlers
Add below worker handlers in `cmd/worker.go` file.
```go
type Handler1 struct {
}

func (w Handler1) Handle(ctx context.Context, payload client.IPayload) error {
//Add your worker code here
return nil
}

func (w Handler1) ZeroPayload() client.IPayload {
return client.NewPayload("", "", "")
}

func NewHandler1() client.JobHandler {
return &Handler1{}
}
```

Register your workers with the topics in `cmd/worker_mapping.go` file.
```go
topicsToHandlersMapping = map[string]client.JobHandler{
		"topic1": NewHandler1(),
	}
```

#### Publish messages
Your server is running on port: 9500, use the below API to publish on a specific topic.
```curl
curl --location 'localhost:9500/producer/publish/topic1' \
--header 'Content-Type: application/json' \
--data '{
    "key":"key1",
    "data":"This is new data"
}'
```
