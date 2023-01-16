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
	ShopID      string    `json:"shopID" gorm:"foreign_key;"`
	Type        string    `json:"type"`
	Species     string    `json:"species"`
	Color       string    `json:"color"`
	Sex         string    `json:"sex"`
	Weight      float32   `json:"weight"`
	Height      float32   `json:"height"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	Age         int16     `json:"age"`
	Tag         string    `json:"tag"`
}

// type Type string

// const (
// 	Dog        Type = "dog"
// 	Cat        Type = "cat"
// 	Fish       Type = "fish"
// 	Rabbit     Type = "rabbit"
// 	Mice       Type = "mice"
// 	Ant        Type = "ant"
// 	Brid       Type = "brid"
// 	Amphibians Type = "amphibians"
// 	Reptiles   Type = "reptiles"
// )

// func (e Type) String() string {
// 	switch e {
// 	case Dog:
// 		return "Dog"
// 	case Cat:
// 		return "Cat"
// 	case Fish:
// 		return "Fish"
// 	case Rabbit:
// 		return "Rabbit"
// 	case Mice:
// 		return "Mice"
// 	case Ant:
// 		return "Ant"
// 	case Brid:
// 		return "Brid"
// 	case amphibians:
// 		return "amphibians"
// 	case reptiles:
// 		return "reptiles"
// 	default:
// 		return fmt.Sprintf("%d", int(e))
// 	}
// }
