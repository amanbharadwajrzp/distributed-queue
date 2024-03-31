package client

type Payload struct {
	Key   string      `json:"key"`
	Topic string      `json:"topic"`
	Data  interface{} `json:"data"`
}

func (p Payload) GetKey() string {
	return p.Key
}

func (p Payload) SetKey(key string) {
	p.Key = key
}

func (p Payload) SetTopic(topicName string) {
	p.Topic = topicName
}

func (p Payload) SetData(data interface{}) {
	p.Data = data
}

func (p Payload) GetTopic() string {
	return p.Topic
}

func (p Payload) GetData() interface{} {
	return p.Data
}

type IPayload interface {
	GetTopic() string
	SetTopic(topicName string)
	GetData() interface{}
	SetData(data interface{})
	GetKey() string
	SetKey(key string)
}

func NewPayload(topic string, key string, data interface{}) IPayload {
	return &Payload{
		Topic: topic,
		Data:  data,
	}
}
