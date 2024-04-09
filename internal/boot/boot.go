package boot

import (
	"distributed-queue/internal/config"
	"distributed-queue/internal/registry"
	"fmt"
	"log"
)

var (
	// Config holds info about growth config
	Config                  config.AppConfig
	ManagerRegistryInstance registry.IManagerRegistry
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
	registry.DaoRegistryInstance = registry.RegisterDao()
	ManagerRegistryInstance = registry.RegisterManagers(registry.DaoRegistryInstance)
}

func GetAppConfig() config.App {
	return Config.App
}
