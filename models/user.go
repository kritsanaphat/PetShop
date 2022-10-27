package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primarykey"`
	Fullname string    `json:"fullanme"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}
