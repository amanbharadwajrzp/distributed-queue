package deserializers

import "distributed-queue/internal/utils"

type IDeSerializable interface {
	DeSerialize() (interface{}, error)
}

func GetDeSerializer(encodedValue []byte, serializableType utils.SerializableType) IDeSerializable {
	switch serializableType {
	/*case utils.JSON:
	return &JSONDeserializer{}*/
	case utils.String:
		return NewStringDeSerializable(encodedValue, serializableType)
	/*case utils.Avro:
	return &AvroDeserializer{}*/
	default:
		return nil
	}
}
