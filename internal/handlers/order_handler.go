package handlers

import (
	"log"
	"net/url"
	"tokos-ws/internal/database"
)

func HandleNewOrder(data interface{}, orderRepo *database.OrderRepository) {
	// Assuming data is a string (the token)
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

	// Now use the decoded token to find the order
	order, err := orderRepo.FindOrderByToken(decodedToken)
	if err != nil {
		log.Printf("Failed to find order by token: %v", err)
		return
	}

	log.Printf("Order found: %+v", order)
	// You can then continue processing the order or broadcast it as needed
}
