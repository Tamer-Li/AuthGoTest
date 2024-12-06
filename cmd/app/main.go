package main

import (
	"AuthGoTest/internal/api"
	"AuthGoTest/internal/config"
	"flag"
	"log"
	"net/http"
)

func main() {
	envStringPath := envFile()

	cfg := config.MustLoad(envStringPath)

	_ = cfg

	srv := api.New()

	http.ListenAndServe("localhost:8080", srv.Router())
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
