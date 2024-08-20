package database

import (
	"database/sql"
	"fmt"
	"log"
	"tokos-ws/internal/database/models"
)

type OrderRepository struct {
	db *sql.DB
}

// Constructor
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) FindOrderByToken(token string) (*models.Order, error) {
	var order models.Order
	// log.Printf("Looking for order with token: %s", token)

	query := `SELECT id, "orderedBy", "isOrderVerified", "status", "verificationToken", "userId", "orderDateTime", "createdAt"
			  FROM "Order" 
			  WHERE "verificationToken" = $1`

	err := repo.db.QueryRow(query, token).Scan(
		&order.ID,
		&order.OrderedBy,
		&order.IsOrderVerified,
		&order.Status,
		&order.VerificationToken,
		&order.UserID,
		&order.OrderDateTime,
		&order.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No order found with token: %s", token)
			return nil, fmt.Errorf("no order found with token: %s", token)
		}
		log.Printf("Error occurred: %v", err)
		return nil, err
	}

	// log.Printf("Order found: %+v", order)
	return &order, nil
}
