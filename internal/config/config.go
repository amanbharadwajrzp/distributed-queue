package config

type AppConfig struct {
	Broker Broker
}

type Broker struct {
	Hostname string
	Port     int
	Topics   []Topic
}

type Topic struct {
	Name       string
	Partitions int
}
