package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
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

func IndexProducts(client *elastic.Client, products []Product) {
	for _, product := range products {
		_, err := client.Index().
			Index("products").
			Id(product.ID.String()).
			BodyJson(product).
			Do(context.Background())
		if err != nil {
			log.Printf("Error to index products %s: %v", product.ID, err)
		}
	}
	log.Println("Products indexed successfully!")
}

func SearchProducts(client *elastic.Client, query string) ([]Product, error) {
	queryBuilder := elastic.NewMultiMatchQuery(query, "name", "description").
		Fuzziness("AUTO").
		PrefixLength(2).
		Operator("AND")

	searchResult, err := client.Search().
		Index("products").
		Query(queryBuilder).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	var products []Product
	for _, hit := range searchResult.Hits.Hits {
		var product Product
		if err := json.Unmarshal(hit.Source, &product); err == nil {
			products = append(products, product)
		}
	}
	return products, nil
}
