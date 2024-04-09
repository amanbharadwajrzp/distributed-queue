package config

type AppConfig struct {
	App App
}

type App struct {
	Hostname string
	Port     int
	Topics   []*Topic
}

type Topic struct {
	Name       string
	Partitions int
}
