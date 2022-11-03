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
	Address  Address   `gorm:"column:Fullname"`
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
	Fullname    string `gorm:"column:fullname"`
	House       string `gorm:"column:house"`
	District    string `gorm:"column:district"`
	Subdistrict string `gorm:"column:subdistrict"`
	City        string `gorm:"column:city"`
	Postcode    string `gorm:"column:postcode"`
}
