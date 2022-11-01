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
	Address  Address   `gorm:"foreignKey:Fullname;references:Fullname"`
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
	gorm.Model
	Fullname    string `gorm:"column:fullname"`
	House       string `gorm:"column:house;default:null"`
	District    string `gorm:"column:district;default:null"`
	Subdistrict string `gorm:"column:subdistrict;default:null"`
	City        string `gorm:"column:city;default:null"`
	Postcode    string `gorm:"column:postcode;default:null"`
}
