package boot

import (
	"distributed-queue/internal/config"
	"fmt"
	"log"
)

var (
	// Config holds info about growth config
	Config config.AppConfig
)

func init() {
	fmt.Println("Initializing App ...", "default")
	Config = initConfig()
}

func initConfig() config.AppConfig {
	err := config.NewDefaultConfig().Load("default", &Config)
	if err != nil {
		log.Fatal(err)
	}
	return Config
}

func Initialize() {
	DaoRegistryInstance = RegisterDao()
	RegisterManagers(DaoRegistryInstance)
}
