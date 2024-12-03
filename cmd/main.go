package main

import (
	"log"
	"net/http"

	"github.com/azevedoguigo/go-search-product-engine.git/config"
	"github.com/azevedoguigo/go-search-product-engine.git/internal"
	"github.com/azevedoguigo/go-search-product-engine.git/pkg"
)

func main() {
	pkg.LoadEnv()

	db := config.InitDB()
	client := config.ConnectElasticsearch()

	products, err := internal.FetchProductsFromDB(db)
	if err != nil {
		log.Fatalf("Error to get products in database: %v", err)
	}
	internal.IndexProducts(client, products)

	http.HandleFunc("/search", internal.HandleSearch(client))
	log.Println("Server running in port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
