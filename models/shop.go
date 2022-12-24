package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	AccountID string    `json:"accountID"`
	ShopID    uuid.UUID `json:"ID" gorm:"primary_key;"`
	ShopName  string    `json:"shopname"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Phone     string    `json:"phone"`
}

type Pet struct {
	gorm.Model
	ID          uuid.UUID `json:"ID" gorm:"primary_key;"`
	ShopID      uuid.UUID `json:"shopID" gorm:"foreign_key;"`
	Type        string    `json:"type"`
	Species     string    `json:"species"`
	Color       string    `json:"color"`
	Sex         string    `json:"sex"`
	Weight      float32   `json:"weight"`
	Height      float32   `json:"height"`
	Price       float32   `json:"price"`
	Description string    `json:"detail"`
	Age         string    `json:"age"`
}
