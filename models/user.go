package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Fullname   string    `json:"fullname"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Address    Address   `gorm:"foreignKey:name;references:fullname"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updtaed_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Address struct {
	gorm.Model
	Name        string `json:"name"`
	House       string `json:"house"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	City        string `json:"city" `
	Postcode    string `json:"postcode"`
}
