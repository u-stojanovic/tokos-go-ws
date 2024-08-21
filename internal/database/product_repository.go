package database

import (
	"database/sql"
	"fmt"
	"log"
	"tokos-ws/internal/database/models"
)

type ProductRepository struct {
	db *sql.DB
}

// Constructor
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) FindProduct(id int) (*models.Product, error) {
	var product models.Product

	query := `SELECT id, name, description, price, category, createdAt 
	          FROM "Product" 
	          WHERE id = $1`

	err := repo.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No product found with ID: %d", id)
			return nil, fmt.Errorf("no product found with ID: %d", id)
		}
		log.Printf("Error occurred while finding product: %v", err)
		return nil, err
	}

	return &product, nil
}
