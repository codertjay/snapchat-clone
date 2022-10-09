package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	snapchat_clone "snapchat-clone/snapchat-clone/database"
	"time"
)

var validate = validator.New()

// UserSignup /* This is meant to register a user*/
func UserSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// initialize a user struct
		var user User
		var foundUser User

		/* Convert json and binds it for golang to understand*/
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// validating
		validatorErr := validate.Struct(user)
		if validatorErr != nil {
			c.JSON(400, gin.H{"error": validatorErr.Error()})
			return
		}

		// Check if the user exists
		db := snapchat_clone.DBConnection()
		db.Where(&User{Email: user.Email}).Find(&foundUser)
		if foundUser.Email != nil {
			c.JSON(400, gin.H{"error": "User With the email already exist"})
			return
		}
		db.Where(&User{Phone: user.Phone}).Find(&foundUser)
		if foundUser.Phone != nil {
			c.JSON(400, gin.H{"error": "User With the Phone number already exist"})
			return
		}

		user, err := CreateUser(&user)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		/* Return the data*/
		c.JSON(201, gin.H{
			"email":        user.Email,
			"token":        user.AccessToken,
			"refreshToken": user.RefreshToken,
		})
		return
	}
}
