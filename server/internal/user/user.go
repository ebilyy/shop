package user

import "gorm.io/gorm"

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

// тимчасово зберігаємо товари в пам'яті
var Users = []User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
}