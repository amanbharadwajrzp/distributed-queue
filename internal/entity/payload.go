package entity

import (
	"distributed-queue/internal/deserializers"
	"distributed-queue/internal/serializers"
	"distributed-queue/internal/utils"
)

type Payload struct {
	Key        string                 `json:"key"`
	Topic      string                 `json:"topic"`
	Data       []byte                 `json:"data"`
	Serializer utils.SerializableType `json:"serializer"`
}

func (p Payload) GetKey() string {
	return p.Key
}

func (p Payload) GetTopic() string {
	return p.Topic
}

func (p Payload) GetData() (interface{}, error) {
	return deserializers.GetDeSerializer(p.Data, p.Serializer).DeSerialize()
}

type IPayload interface {
	GetTopic() string
	GetData() (interface{}, error)
	GetKey() string
}

func NewPayload(topic string, key string, data interface{}, serializer utils.SerializableType) IPayload {
	var encodedValue []byte
	if data != nil {
		encodedValue, _ = serializers.GetSerializer(data, serializer).Serialize()
	}
	return &Payload{
		Topic:      topic,
		Key:        key,
		Data:       encodedValue,
		Serializer: serializer,
	}
}
