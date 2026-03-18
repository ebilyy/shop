package user

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Password string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName string  `json:"last_name"`
	Phone string  `json:"phone"`
	Address string  `json:"address"`
	City string  `json:"city"`
	Country string  `json:"country"`
	AcceptedTerms bool  `json:"accepted_terms"`
	Comment string  `json:"comment"`
	RegistrationDate string `json:"registration_date"`
	IsDeleted bool  `json:"is_deleted"`
	Surname string  `json:"surname"`
	Status string  `json:"status"`
	Avatar string  `json:"avatar"`
}

// // тимчасово зберігаємо товари в пам'яті
var Users = []User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
}
type Handler struct { DB *gorm.DB }

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user := r.Body
	fmt.Println(user)
	return
}

// func HandleUsersGet(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 			return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	users := h.DB.Find(&User{})
// 	fmt.Println(users)
// 	if err := json.NewEncoder(w).Encode(Users); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 	}
// 	// if err := json.NewEncoder(w).Encode(app.db.Find(&User{})); err != nil {
// 	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 		return
// 	// }
// }