package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func UserMiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //valid
			fmt.Println("FROM : Middleware UserID", claims["ID"])
			c.Set("AccountID", claims["ID"])
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "error", "messege": err.Error()})
			return
		}

		// before request
		c.Next()
	}
}

func ShopMiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY2"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //valid
			fmt.Println("FROM : Middleware ShopID", claims["ID"])
			c.Set("ShopID", claims["ID"])
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "error", "messege": err.Error()})
			return
		}

		// before request
		c.Next()
	}
}

func AdminMiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY3"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //valid
			fmt.Println("FROM : Middleware AddminID", claims["ID"])
			c.Set("AdminID", claims["ID"])
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "error", "messege": err.Error()})
			return
		}

		// before request
		c.Next()
	}
}
