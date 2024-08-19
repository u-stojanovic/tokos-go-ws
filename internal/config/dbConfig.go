package config

import (
	"os"
)

type Config struct {
	DatabaseDSN string
	ServerPort  string
}

func LoadConfig() Config {
	return Config{
		DatabaseDSN: os.Getenv("DATABASE_URL"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}
}
