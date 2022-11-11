package models

import (
	"database/sql/driver"

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
	PetType petType   `gorm:"type:pet_type"`
	Species string    `json:"species"`
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

type petType string

const (
	DOG    petType = "DOG"
	CAT    petType = "CAT"
	RABBIT petType = "RABBIT"
	FISH   petType = "FISH"
)

func (ct *petType) Scan(value interface{}) error {
	*ct = petType(value.([]byte))
	return nil
}

func (ct petType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (Pet) TableName() string {
	return "my_table"
}
