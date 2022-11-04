package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Price  int32     `json:"price"`
	Detail string    `json:"detail"`
}

type Shop struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"primary_key;"`
	Nameshop string    `json:"fullname"`

	// Address  Address   `gorm:"column:Fullname"`
}
