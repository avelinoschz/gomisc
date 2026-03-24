package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CatalogStore struct {
	products map[string]Product
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{
		products: map[string]Product{
			"HAMMER-001": {SKU: "HAMMER-001", Name: "Hammer", Price: 25},
		},
	}
}

func (s *CatalogStore) Save(product Product) error {
	// TODO: implement
	return nil
}

func main() {
	store := NewCatalogStore()

	http.HandleFunc("POST /catalog", func(w http.ResponseWriter, r *http.Request) {
		_ = store
		_ = json.NewDecoder
		// TODO: implement
	})

	_ = http.ListenAndServe(":8080", nil)
}
