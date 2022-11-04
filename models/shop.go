package models

import (
	"github.com/gofrs/uuid"
)

type Shop struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Price  int32     `json:"price"`
	Detail string    `json:"detail"`
}
