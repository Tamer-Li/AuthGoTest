package config

import "time"

type Config struct {
	Env        string `yaml:"env" env-default:"test"`
	UrlDB      `yaml:"url_db"`
	HTTPServer `yaml:"http_server"`
	JWTSecret  `yaml:"jwt_secret"`
}

type UrlDB struct {
	DBHost string `yaml:"db_host"`
	DBPort string `yaml:"db_port"`
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
	DBName string `yaml:"db_name"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type JWTSecret struct {
	Secret string `yaml:"secret" env-default:"secret"`
}
