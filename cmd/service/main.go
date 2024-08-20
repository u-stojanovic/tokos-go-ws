package main

import (
	"log"
	"net/http"
	"tokos-ws/internal/config"
	"tokos-ws/internal/database"
	"tokos-ws/internal/websocket"

	"github.com/joho/godotenv"
)

func main() {
	// .env load
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// config load
	cfg := config.LoadConfig()

	// db init
	db, err := database.InitDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close() // db conn closing after app shut down

	// orderRepo init
	orderRepo := database.NewOrderRepository(db)

	// orderRepo passed to ws handler
	websocket.SetOrderRepository(orderRepo)

	// starting broadcasting
	go websocket.Start()

	// handling ws requests
	http.HandleFunc("/ws", websocket.HandleConnections)

	// starting the server
	log.Println("WebSocket server started on port:", cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
