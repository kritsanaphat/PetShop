package models

import (
	"database/sql/driver"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Pet struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Petype Petype    `json:"pet_type"`
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

type Petype string

const (
	DOG Petype = "DOG"
	CAT Petype = "CAT"
)

func (ct *Petype) Scan(value interface{}) error {
	*ct = Petype(value.([]byte))
	return nil
}

func (ct Petype) Value() (driver.Value, error) {
	return string(ct), nil
}
