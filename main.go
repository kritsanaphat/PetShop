package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	// router := mux.NewRouter()
	// p := models.User{Fullname: "krit"}
	// fmt.Print(p)
	r := gin.New()
	r.GET("/user", h.GetRegister)
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	http.ListenAndServe(":8080", r)
}
