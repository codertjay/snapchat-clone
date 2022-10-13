package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"snapchat-clone/models"
	"snapchat-clone/snapchat-clone/database"
	"snapchat-clone/utils"

	"time"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.DBConnection()
		defer database.CloseDB()
		/* get the token*/
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(400, gin.H{"error": "Authorization token not passed"})
			c.Abort()
			return
		}
		claims, err := ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		var user models.User
		db.First(&user, claims.ID)
		c.Set("user", user)
		c.Next()

	}
}

func ValidateToken(signedToken string) (claims *utils.SignedInDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&utils.SignedInDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*utils.SignedInDetails)
	if !ok {
		msg = fmt.Sprint("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
