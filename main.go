package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	DB := db.Init()
	h := handlers.New(DB)
	// router := mux.NewRouter()
	// p := models.User{Fullname: "krit"}
	// fmt.Print(p)
	r := gin.New()
	r.Use(CORSMiddleware())
	r.GET("/user", h.GetRegister)
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/addpet", h.AddPet)
	http.ListenAndServe(":8080", r)
}
