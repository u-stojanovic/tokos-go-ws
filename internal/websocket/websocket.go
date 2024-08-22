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
	upgrader    = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	orderRepo   *database.OrderRepository
	productRepo *database.ProductRepository
)

// SetOrderRepository sets the OrderRepository that will be used in handlers.
func SetOrderRepository(repo *database.OrderRepository) {
	orderRepo = repo
}

func SetProductRepository(repo *database.ProductRepository) {
	productRepo = repo
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "https://ataoakaoasa.vercel.app, https://ataoakaoasadmin-panel.vercel.app, http://localhost:3001, http://localhost:3000, http://localhost:8000, https://tokos-go-ws-production.up.railway.app")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	// Upgrading to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer func() {
		broadcast.RemoveClient(ws)
		ws.Close()
		log.Println("WebSocket connection closed:", ws.RemoteAddr())
	}()

	broadcast.RegisterClient(ws)
	log.Println("WebSocket connection established with client:", ws.RemoteAddr())

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
			handlers.HandleNewProduct(msg.Data, productRepo)
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
