package serializers

import "distributed-queue/internal/utils"

type ISerializable interface {
	Serialize() ([]byte, error)
	Get() interface{}
}

func GetSerializer(value interface{}, serializableType utils.SerializableType) ISerializable {
	switch serializableType {
	/*case utils.JSON:
	return &JSONDeserializer{}*/
	case utils.String:
		return NewStringSerializable(value)
	/*case utils.Avro:
	return &AvroDeserializer{}*/
	default:
		return nil
	}
}
