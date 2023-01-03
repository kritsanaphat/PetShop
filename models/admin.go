package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	AdminID    uuid.UUID `json:"accountID" gorm:"primary_key;"`
	EmployeeID string    `json:"employeeID"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
}
