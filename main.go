package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
	"github.com/kritsanaphat/PetShop/middleware"
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
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := db.Init()
	h := handlers.New(DB)

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/getUserByID/:ID", h.Profile)
	r.PATCH("/updateUser/:ID", h.UpdateAddress)
	r.POST("/register", h.UserRegister)
	r.POST("/login", h.UserLogin)
	r.POST("/addpet", h.AddPet)
	r.GET("/allpet", h.GetAllPet)

	authorized := r.Group("/user", middleware.MiddlewareJWT())
	authorized.GET("getAllUser", h.GetAllUser)
	authorized.GET("/profile", h.Profile)

	http.ListenAndServe(":8080", r)
}
