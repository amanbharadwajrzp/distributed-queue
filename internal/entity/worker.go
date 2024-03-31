package entity

type Worker struct {
	workerId string
	broker   IBroker
}

type IWorker interface {
	GetWorkerId() string
	GetBroker() IBroker
}

func NewWorker(workerId string, broker IBroker) IWorker {
	return &Worker{
		workerId: workerId,
		broker:   broker,
	}
}

func (w Worker) GetWorkerId() string {
	return w.workerId
}

func (w Worker) GetBroker() IBroker {
	return w.broker
}
