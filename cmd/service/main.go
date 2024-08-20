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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load Config
	cfg := config.LoadConfig()

	// Initialize the database connection
	db, err := database.InitDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when the application shuts down

	// Initialize OrderRepository
	orderRepo := database.NewOrderRepository(db)

	// Pass the repository to the WebSocket handler
	websocket.SetOrderRepository(orderRepo)

	// Handle WebSocket requests
	http.HandleFunc("/ws", websocket.HandleConnections)

	// Start the server on the specified port
	log.Println("WebSocket server started on port:", cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
