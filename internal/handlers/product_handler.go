package handlers

import "log"

// HandleNewProduct processes the new product data and broadcasts it to all clients.
func HandleNewProduct(data interface{}) {
	// broadcast.Broadcast(broadcast.Message{Event: "update_product_list", Data: data})
	log.Printf("Processing new product: %v", data)
}
