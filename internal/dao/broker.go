package dao

import "distributed-queue/internal/entity"

type Broker struct {
	brokers               []entity.IBroker
	brokerToTopicsMapping map[string][]string
}

func (b Broker) Get(brokerId string) entity.IBroker {
	for _, broker := range b.brokers {
		if entity.GetBrokerId(broker.GetHost(), broker.GetPort()) == brokerId {
			return broker
		}
	}
	return nil
}

func (b Broker) Create(broker entity.IBroker) bool {
	if b.brokers == nil || len(b.brokers) == 0 {
		b.brokers = make([]entity.IBroker, 0)
	}
	b.brokers = append(b.brokers, broker)
	return true
}

func (b Broker) AddTopic(brokerId string, topic string) bool {
	if b.brokerToTopicsMapping[brokerId] == nil || len(b.brokerToTopicsMapping[brokerId]) == 0 {
		b.brokerToTopicsMapping[brokerId] = make([]string, 0)
	}
	b.brokerToTopicsMapping[brokerId] = append(b.brokerToTopicsMapping[brokerId], topic)
	return true
}

func (b Broker) DeleteTopic(brokerId string, topics string) bool {
	if b.brokerToTopicsMapping[brokerId] == nil || len(b.brokerToTopicsMapping[brokerId]) == 0 {
		return false
	}
	b.brokerToTopicsMapping[brokerId] = nil
	return true
}

func (b Broker) List() []entity.IBroker {
	return b.brokers
}

func (b Broker) GetTopics(brokerId string) []string {
	return b.brokerToTopicsMapping[brokerId]
}

type IBroker interface {
	Create(broker entity.IBroker) bool
	AddTopic(brokerId string, topics string) bool
	DeleteTopic(brokerId string, topics string) bool
	List() []entity.IBroker
	Get(brokerId string) entity.IBroker
	GetTopics(brokerId string) []string
}

func NewBroker() IBroker {
	return &Broker{
		brokers:               make([]entity.IBroker, 0),
		brokerToTopicsMapping: make(map[string][]string),
	}
}
