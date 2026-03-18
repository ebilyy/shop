package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"schemka.store/app/internal/config"
	"schemka.store/app/internal/user"
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

type application struct {
	db *gorm.DB
}

func openDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (app *application) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var u user.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := app.db.Create(&u).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	case http.MethodGet:
		var users []user.User
		if err := app.db.Find(&users).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) runMigrations() {
	if err := app.db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("auto-migrate users failed: %v", err)
	}
}

func main() {
	cfg := config.Load()

	db, err := openDB(cfg.DBDSN)

	// if err := db.AutoMigrate(&user.User{}); err != nil {
	// 	log.Fatalf("auto-migrate failed: %v", err)
	// }

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	app := &application{db: db}
	app.runMigrations()

	mux := http.NewServeMux()

	// GET /products – список товарів
	// mux.HandleFunc("/products", handleProducts)
	mux.HandleFunc("/users", app.handleUsers)
	addr := ":8081"
	log.Printf("Starting server on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(products); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// }
}
