package repositories

import (
	"github.com/kritsanaphat/PetShop/models"
	"gorm.io/gorm"
)

type todoRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) todoRepository {
	return todoRepository{db}
}

func (h todoRepository) UserRegister(todo *models.Account) (err error) {
	if err = h.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
