package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"snapchat-clone/snapchat-clone/helpers"
	"time"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* get the token*/
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(400, gin.H{"error": "Authorization token not passed"})
			c.Abort()
			return
		}
		claims, err := ValidateToken(clientToken)
		if err != "" {
			c.JSON(500, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("name ", claims.Name)
		c.Set("user_id", claims.ID)
		c.Next()

	}
}

func ValidateToken(signedToken string) (claims *helpers.SignedInDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&helpers.SignedInDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*helpers.SignedInDetails)
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
