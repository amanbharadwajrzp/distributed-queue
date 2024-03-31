package serializers

import "distributed-queue/internal/utils"

type StringSerializable struct {
	Value interface{}            `json:"value"`
	Type  utils.SerializableType `json:"type"`
}

func NewStringSerializable(value interface{}, t utils.SerializableType) ISerializable {
	return &StringSerializable{
		Value: value,
		Type:  t,
	}
}

func (s StringSerializable) Serialize() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
