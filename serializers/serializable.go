package serializers

import "mcoding/utils"

type ISerializable interface {
	Serialize() ([]byte, error)
}

func GetSerializer(value interface{}, serializableType utils.SerializableType) ISerializable {
	switch serializableType {
	/*case utils.JSON:
	return &JSONDeserializer{}*/
	case utils.String:
		return NewStringSerializable(value, serializableType)
	/*case utils.Avro:
	return &AvroDeserializer{}*/
	default:
		return nil
	}
}
