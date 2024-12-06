package config

import (
	"log"
	"os"
)

type Config struct {
	UrlDB `yaml:"url_db" defa`
}

type UrlDB struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func MustLoad(configPath string) *Config {
	cfgPath := os.Getenv(configPath)
	if cfgPath == "" {
		log.Fatal("CONFIG PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cfgPath)
	}

	env
}
