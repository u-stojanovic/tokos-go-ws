package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DatabaseDSN string
	ServerPort  string
}

func LoadConfig() Config {
	databaseDSN := os.Getenv("DATABASE_URL")
	if databaseDSN == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	fmt.Println("serverPort: ", serverPort)

	return Config{
		DatabaseDSN: databaseDSN,
		ServerPort:  serverPort,
	}
}
