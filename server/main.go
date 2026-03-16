package main

import (
		"encoding/json"
		"log"
		"net/http"
)

type Product struct {
		ID    int64   `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
}

// тимчасово зберігаємо товари в пам'яті
var products = []Product{
		{ID: 1, Name: "Engine Oil", Price: 29.99},
		{ID: 2, Name: "Brake Pads", Price: 49.99},
}

func main() {
		mux := http.NewServeMux()

		// GET /products – список товарів
		mux.HandleFunc("/products", handleProducts)

		// TODO: додати інші роутери: /products/{id}, /cart, /orders тощо

		addr := ":8081"
		log.Printf("Starting server on %s\n", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
				log.Fatalf("server error: %v", err)
		}
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
}