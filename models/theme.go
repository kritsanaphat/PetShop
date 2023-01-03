package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Theme struct {
	gorm.Model
	ThemeID  uuid.UUID `json:"themeID" gorm:"primary_key;"`
	AuthorID string    `json:"authorID"`
	Topic    string    `json:"topic"`
	Content  string    `json:"content"`
}

type Comment struct {
	gorm.Model
	AuthorID  string    `json:"authorID"`
	ThemeID   uuid.UUID `json:"themeID" gorm:"foreign_key;"`
	CommentID uuid.UUID `json:"commentID" gorm:"primary_key;"`
	Comment   string    `json:"comment"`
}
