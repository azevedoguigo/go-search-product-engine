package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/olivere/elastic/v7"
)

func HandleSearch(client *elastic.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Query not provided", http.StatusBadRequest)
			return
		}

		products, err := SearchProducts(client, query)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error to get products: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
