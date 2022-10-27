package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Fullname string    `json:"fullname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

func (base *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV3())
}
