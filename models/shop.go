package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Type   petType   `gorm:"type:ENUM('DOG', 'CAT')"`
	Price  int32     `json:"price"`
	Detail string    `json:"detail"`
}

type Shop struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"primary_key;"`
	Nameshop string    `json:"fullname"`
	Phone    string    `json:"phone"`

	// Address  Address   `gorm:"column:Fullname"`
}

type petType string

const (
	DOG petType = "DOG"
	CAT petType = "CAT"
)
