package deserializers

import (
	"distributed-queue/internal/utils"
)

type StringDeserializer struct {
	EncodedValue []byte                 `json:"encoded_value"`
	Type         utils.SerializableType `json:"type"`
}

func NewStringDeSerializable(encodedValue []byte, t utils.SerializableType) IDeSerializable {
	return &StringDeserializer{
		EncodedValue: encodedValue,
		Type:         t,
	}
}

func (s StringDeserializer) DeSerialize() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
