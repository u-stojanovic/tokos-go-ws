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
	// Upgrading to ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer func() {
		// client is removed when the connection is closed
		broadcast.RemoveClient(ws)
		ws.Close()
		log.Println("WebSocket connection closed:", ws.RemoteAddr())
	}()

	// client is registered once the connection is established
	broadcast.RegisterClient(ws)
	log.Println("WebSocket connection established with client:", ws.RemoteAddr())

	// Infinite loop to keep the connection open
	for {
		var msg broadcast.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket Read Error: %v", err)
			}
			break
		}

		switch msg.Event {
		case "new_product":
			handlers.HandleNewProduct(msg.Data)
		case "new_order":
			handlers.HandleNewOrder(msg.Data, orderRepo)
		default:
			log.Println("Unknown event type:", msg.Event)
		}
	}
}

func Start() {
	log.Println("StartBroadcasting called in start")
	go broadcast.StartBroadcasting()
}
