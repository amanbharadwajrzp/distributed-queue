package serializers

import "encoding/json"

type StringSerializable struct {
	Input  interface{} `json:"input"`
	Output string      `json:"output"`
}

func NewStringSerializable(input interface{}) ISerializable {
	return &StringSerializable{
		Input: input,
	}
}

func (s StringSerializable) Serialize() ([]byte, error) {
	output, err := json.Marshal(s.Input)
	if err != nil {
		return nil, err
	}
	s.Output = string(output)
	return output, nil
}

func (s StringSerializable) Get() interface{} {
	return s.Output
}
