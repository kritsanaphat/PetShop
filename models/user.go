package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Fullname string    `json:"fullname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
