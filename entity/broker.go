package entity

type Broker struct {
	host string `json:"host"`
	port int    `json:"port"`
}

type IBroker interface {
	GetHost() string
	GetPort() int
}

func NewBroker(host string, port int) IBroker {
	return &Broker{
		host: host,
		port: port,
	}
}

func (b *Broker) GetHost() string {
	return b.host
}

func (b *Broker) GetPort() int {
	return b.port
}

func GetBrokerId(host string, port int) string {
	return host + ":" + string(rune(port))
}
