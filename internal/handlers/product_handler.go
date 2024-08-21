package handlers

import (
	"log"
	"tokos-ws/internal/broadcast"
	"tokos-ws/internal/database"
)

func HandleNewProduct(data interface{}, productRepo *database.ProductRepository) {
	id, ok := data.(int)
	if !ok {
		log.Println("Invalid product ID format")
		return
	}

	product, err := productRepo.FindProduct(id)
	if err != nil {
		log.Printf("Failed to find product by ID: %v", err)
		return
	}

	log.Printf("Product found: %+v", product)

	message := broadcast.Message{
		Event: "new_product",
		Data:  product,
	}
	broadcast.Broadcast(message)
}
