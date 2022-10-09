package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
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
		snapchat_clone.CloseDB()
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
			"email":         user.Email,
			"access_token":  user.AccessToken,
			"refresh_token": user.RefreshToken,
		})
		return
	}
}

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		db := snapchat_clone.DBConnection()
		snapchat_clone.CloseDB()
		/*  Initialize the user*/
		var user User
		var foundUser User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid parameters pass"})
			return
		}
		// check if the email is passed
		if user.Email == nil {
			c.JSON(400, gin.H{"error": "Email not passed"})
			return
		}
		if user.Password == nil {
			c.JSON(400, gin.H{"error": "Password not passed"})
			return
		}
		//	 check if the user email exists
		db.Where(&User{Email: user.Email}).Find(&foundUser)
		if foundUser.Email == nil {
			c.JSON(400, gin.H{"error": "user does not exist or invalid parameters passed. Please make sure you pass the correct params"})
			return
		}
		check, err := VerifyPassword(*foundUser.Password, *user.Password)
		if err != nil && check == false {
			c.JSON(400, gin.H{"error": "Invalid password or mail passed"})
			return
		}
		token, refreshToken, err := GenerateAllToken(*foundUser.Name, *foundUser.Email, *foundUser.Phone, foundUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error occurred"})
			return
		}

		// create the update sql query
		sqlStatement := `UPDATE users SET access_token=$2, refresh_token=$3 WHERE email=$1`
		// Update the user token
		// execute the sql statement
		err = db.Exec(sqlStatement, foundUser.Email, token, refreshToken).Error
		if err != nil {
			log.Panicln("Error occurred", err)
		}
		c.JSON(200, gin.H{
			"email":         user.Email,
			"access_token":  token,
			"refresh_token": refreshToken,
		})
		return

	}
}
