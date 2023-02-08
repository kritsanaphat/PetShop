package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	AccountID uuid.UUID `json:"accountID" gorm:"primary_key;"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

type Address struct {
	gorm.Model
	ID          uuid.UUID `json:"accountID" gorm:"primary_key"`
	House       string    `json:"house"`
	District    string    `json:"district"`
	Subdistrict string    `json:"subdistrict"`
	City        string    `json:"city"`
	Postcode    string    `json:"postcode"`
}

type Chart struct {
	gorm.Model
	ChartID uuid.UUID `json:"chartID" gorm:"primary_key"`
	ID      string    `json:"accountID"`
	Item    uuid.UUID `json:"itemID"`
}
type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Login struct {
	Username string `json:"username"`
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
