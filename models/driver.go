package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	AccountID uuid.UUID `json:"accountID" gorm:"primary_key;"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

type DriverAddress struct {
	gorm.Model
	ID          uuid.UUID `json:"accountID" gorm:"primary_key"`
	House       string    `json:"house"`
	District    string    `json:"district"`
	Subdistrict string    `json:"subdistrict"`
	City        string    `json:"city"`
	Postcode    string    `json:"postcode"`
}
