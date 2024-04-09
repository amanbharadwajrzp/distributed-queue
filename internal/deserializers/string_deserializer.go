package deserializers

import "encoding/json"

type StringDeserializer struct {
	EncodedValue []byte `json:"encoded_value"`
}

func NewStringDeSerializable(encodedValue []byte) IDeSerializable {
	return &StringDeserializer{
		EncodedValue: encodedValue,
	}
}

func (s StringDeserializer) DeSerialize() (interface{}, error) {
	var outputVal interface{}
	if err := json.Unmarshal(s.EncodedValue, &outputVal); err != nil {
		return nil, err
	}
	return outputVal, nil
}
