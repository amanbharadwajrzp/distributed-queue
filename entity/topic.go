package entity

import "mcoding/utils"

type topic struct {
	id                 string  `json:"id"`
	broker             IBroker `json:"broker"`
	name               string  `json:"name"`
	numberOfPartitions int     `json:"partitions"`
}

func (t topic) SetBroker(broker IBroker) {
	t.broker = broker
}

func (t topic) GetID() string {
	return t.id
}

func (t topic) GetName() string {
	return t.name
}

func (t topic) GetNumberOfPartitions() int {
	return t.numberOfPartitions
}

type Topic interface {
	GetID() string
	GetName() string
	GetNumberOfPartitions() int
	SetBroker(broker IBroker)
}

func NewTopic(name string, numberOfPartitions int) Topic {
	return &topic{
		id:                 utils.NewUUIDAsString(),
		name:               name,
		numberOfPartitions: numberOfPartitions,
	}
}
