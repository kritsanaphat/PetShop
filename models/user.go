package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Fullname string    `json:"fullname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
