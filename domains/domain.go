package domains

import (
	"github.com/kritsanaphat/PetShop/models"
)

type ToDoUseCase interface {
	UserRegister(t *models.Account) (err error)
}

// ToDoRepository ...
type ToDoRepository interface {
	UserRegister(t *models.Account) (err error)
}
