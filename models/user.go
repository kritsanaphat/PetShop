package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountID uuid.UUID `json:"ID" gorm:"primary_key;"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   Address   `gorm:"foreignKey:ID;references:AccountID"`
}

type Address struct {
	ID          uuid.UUID `json:"AddressID" gorm:"primary_key;"`
	Fullname    string    `json:"fullname"`
	House       string    `json:"house"`
	District    string    `json:"district"`
	Subdistrict string    `json:"subdistrict"`
	City        string    `json:"city"`
	Postcode    string    `json:"postcode"`
}

type Register struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Search struct {
	ID uuid.UUID
}

type Favorite struct {
	gorm.Model
	User_ID uuid.UUID
	Pet_ID  uuid.UUID
}
