package websocket

import (
	"log"
	"net/http"
	"tokos-ws/internal/broadcast"
	"tokos-ws/internal/database"
	"tokos-ws/internal/handlers"

	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	orderRepo *database.OrderRepository
)

// SetOrderRepository sets the OrderRepository that will be used in handlers.
func SetOrderRepository(repo *database.OrderRepository) {
	orderRepo = repo
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer ws.Close()

	broadcast.RegisterClient(ws)

	for {
		var msg broadcast.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket Read Error: %v", err)
			broadcast.RemoveClient(ws)
			break
		}

		// Process the message based on the event type
		switch msg.Event {
		case "new_product":
			handlers.HandleNewProduct(msg.Data)
		case "new_order":
			// Use the repository within the handler
			handlers.HandleNewOrder(msg.Data, orderRepo)
		default:
			log.Println("Unknown event type:", msg.Event)
		}
	}
}

func Start() {
	go broadcast.StartBroadcasting()
}
