package deliveries

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/kritsanaphat/PetShop/domains"
	"github.com/kritsanaphat/PetShop/models"
	"golang.org/x/crypto/bcrypt"
)

type ToDoHandler struct {
	todoUseCase domains.ToDoUseCase
}

// NewToDoHandler ...
func NewToDoHandler(usecase domains.ToDoUseCase) *ToDoHandler {
	return &ToDoHandler{
		todoUseCase: usecase,
	}
}

func (t *ToDoHandler) UserRegister(c *gin.Context) {
	var json models.Register

	if err := c.ShouldBindJSON(&json); err != nil { //Check the integrity of the information
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := models.Account{
		AccountID: uuid,
		Password:  string(encryptedPassword),
		Username:  json.Username,
		Email:     json.Email,
		Phone:     json.Phone,
	}

	err = t.todoUseCase.UserRegister(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "error", "messege": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}
