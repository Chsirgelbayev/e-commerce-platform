package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	isDebug       bool `env: "IS_DEBUG" env-default: "false"`
	isDevelopment bool `env: "IS_DEV" env-default: "false"`
	Listen        struct {
		Type   string `env: "TYPE" env-default: "port"`
		BindIP string `env: "BIND_IP" env-default: "0.0.0.0"`
		Port   string `env: "PORT" env-default: "8080"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env: "ADMIN_EMAIL" env-required: "true"`
			Password string `env: "ADMIN_PASSWORD" env-required: "true"`
		}
	}
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		config = &Config{}
		cleanenv.ReadEnv(config)

		if err := cleanenv.ReadEnv(config); err != nil {
			descriptionText := "config validation failed"
			helpText, _ := cleanenv.GetDescription(config, &descriptionText)

			log.Print(helpText)
			log.Fatal(err)
		}
	})
	return config
}
