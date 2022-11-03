package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Fullname string    `gorm:"column:fullname"`
	Password string    `gorm:"column:password"`
	Email    string    `gorm:"column:email"`
	Address  Address   `json:"address"`
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
type Address struct {
	Fullname    string `json:"fullname"`
	House       string `json:"house"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
}
