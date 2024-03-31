# Distrubuted In-memory Queue.
A simple in-memory queue that supports partitioning and offset management. Multiple consumer workers can subscribe to a single topic.

#### Local Setup Setup broker and worker servers.

- Run `make go-build-broker` to build api and migrations docker image.
- Run `make go-build-worker` to build api docker image.

#### Register worker handlers
Add below worker handlers in `e2e/worker.go` file.
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

Register your workers with the topics in `e2e/worker_mapping.go` file.
```go
topicsToHandlersMapping = map[string]client.JobHandler{
		"topic1": NewHandler1(),
	}
```
#### Start the services with below commands.