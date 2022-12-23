package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kritsanaphat/PetShop/db"
	"github.com/kritsanaphat/PetShop/handlers"
	"github.com/kritsanaphat/PetShop/middleware"
	_ "github.com/kritsanaphat/PetShop/middleware"
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
	r := gin.New()
	r.Use(CORSMiddleware())

	r.GET("/getAllUser", h.GetAllUser)
	r.GET("/getUserByID/:ID", h.GetUserByID)
	//r.PATCH("/updateUser/:ID", h.UpdateAddress)
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/addpet", h.AddPet)
	r.GET("/allpet", h.GetAllPet)

	authorized := r.Group("/handlers", middleware.MiddlewareJWT)
	authorized.GET("getAllUser", h.GetAllUser)

	http.ListenAndServe(":8080", r)
}
