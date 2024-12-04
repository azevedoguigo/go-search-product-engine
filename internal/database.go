package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       *float64  `json:"price,omitempty"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func FetchProductsFromDB(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
