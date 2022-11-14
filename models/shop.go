package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// type petType string

// const (
// 	DOG petType = "DOG"
// 	CAT petType = "CAT"
// )

type Pet struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	PetType string    `json:"type"`
	Species string    `json:"species"`
	Color   string    `json:"color"`
	Price   int32     `json:"price"`
	Detail  string    `json:"detail"`
}

type Shop struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"primary_key;"`
	Nameshop string    `json:"fullname"`
	Phone    string    `json:"phone"`

	// Address  Address   `gorm:"column:Fullname"`
}
