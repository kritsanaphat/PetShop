package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kritsanaphat/PetShop/deliveries"
	"github.com/kritsanaphat/PetShop/repositories"
	"github.com/kritsanaphat/PetShop/usecases"
)

func SetupRouter() *gin.Engine {

	todoRepo := repositories.NewToDoRepository(db.)
	todoUseCase := usecases.NewToDoUseCase(todoRepo)
	todoHandler := deliveries.NewToDoHandler(todoUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("register", todoHandler.UserRegister)
	}
	return r
}
