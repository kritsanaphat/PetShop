package usecases

import (
	"github.com/kritsanaphat/PetShop/domains"
	"github.com/kritsanaphat/PetShop/models"
)

type todoUseCase struct {
	todoRepo domains.ToDoRepository
}

func NewToDoUseCase(repo domains.ToDoRepository) domains.ToDoUseCase {
	return &todoUseCase{
		todoRepo: repo,
	}
}

func (t *todoUseCase) UserRegister(input *models.Account) (err error) {
	handleErr := t.todoRepo.UserRegister(input)
	return handleErr
}

func (t *todoUseCase) ShopRegister(input *models.Shop) (err error) {
	handleErr := t.todoRepo.ShopRegister(input)
	return handleErr
}
