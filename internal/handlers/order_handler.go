package handlers

import (
	"log"
	"net/url"
	"tokos-ws/internal/broadcast"
	"tokos-ws/internal/database"
)

func HandleNewOrder(data interface{}, orderRepo *database.OrderRepository) {
	token, ok := data.(string)
	if !ok {
		log.Println("Invalid token format")
		return
	}

	// Decode the token from URL encoding
	decodedToken, err := url.QueryUnescape(token)
	if err != nil {
		log.Printf("Failed to decode token: %v", err)
		return
	}

	order, err := orderRepo.FindOrderByToken(decodedToken)
	if err != nil {
		log.Printf("Failed to find order by token: %v", err)
		return
	}

	log.Printf("Order found: %+v", order)

	message := broadcast.Message{
		Event: "new_order",
		Data:  order,
	}
	broadcast.Broadcast(message)
}
