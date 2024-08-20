package broadcast

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	mutex     = &sync.Mutex{}
)

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func Broadcast(message Message) {
	log.Println("message: ", message)
	broadcast <- message
}

func RegisterClient(ws *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	clients[ws] = true
	log.Println("Client registered, total clients:", len(clients))
}

func RemoveClient(ws *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, ws)
}

func StartBroadcasting() {
	for {
		message := <-broadcast
		mutex.Lock()
		log.Println("Broadcasting message to client")
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
