package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Fullname string    `json:"fullname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
