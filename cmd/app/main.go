package main

import (
	"AuthGoTest/internal/config"
	"flag"
	"log"
)

func main() {
	envStringPath := envFile()

	cfg := config.MustLoad(envStringPath)
}

func envFile() string {
	envPath := flag.String(
		"env",
		".env",
		"path to .env file",
	)
	flag.Parse()

	if *envPath == "" {
		log.Fatal("env file not specified")
	}

	return *envPath
}
