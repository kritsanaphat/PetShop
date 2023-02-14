package repositories

import (
	"github.com/kritsanaphat/PetShop/models"
	"gorm.io/gorm"
)

type todoRepository struct {
	conn *gorm.DB
}

func NewToDoRepository(db *gorm.DB) todoRepository {
	return todoRepository{db}
}

func (t todoRepository) UserRegister(todo *models.Account) (err error) {
	if err = t.conn.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (t todoRepository) ShopRegister(todo *models.Shop) (err error) {
	if err = t.conn.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}
