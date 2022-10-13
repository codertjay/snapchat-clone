package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"snapchat-clone/models"
	snapchat_clone "snapchat-clone/snapchat-clone/database"
	"snapchat-clone/utils"
	"time"
)

var validate = validator.New()

// UserSignup /* This is meant to register a user*/
func UserSignup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// initialize a user struct
		var user models.User
		var foundUser models.User

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
		defer snapchat_clone.CloseDB()
		db.Where(&models.User{Email: user.Email}).Find(&foundUser)
		if foundUser.Email != nil {
			c.JSON(400, gin.H{"error": "User With the email already exist"})
			return
		}
		db.Where(&models.User{Phone: user.Phone}).Find(&foundUser)
		if foundUser.Phone != nil {
			c.JSON(400, gin.H{"error": "User With the Phone number already exist"})
			return
		}

		user, err := utils.CreateUser(&user)

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
		defer snapchat_clone.CloseDB()
		/*  Initialize the user*/
		var user models.User
		var foundUser models.User
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
		db.Where(&models.User{Email: user.Email}).Find(&foundUser)
		if foundUser.Email == nil {
			c.JSON(400, gin.H{"error": "user does not exist or invalid parameters passed. Please make sure you pass the correct params"})
			return
		}
		check, err := utils.VerifyPassword(*foundUser.Password, *user.Password)
		if err != nil && check == false {
			c.JSON(400, gin.H{"error": "Invalid password or mail passed"})
			return
		}
		token, refreshToken, err := utils.GenerateAllToken(*foundUser.Name, *foundUser.Email, *foundUser.Phone, foundUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error occurred"})
			return
		}

		// Update the user token
		err = db.Model(&user).Where(&models.User{Email: user.Email}).
			Update("access_token", token).
			Update("refresh_token", refreshToken).Find(&foundUser).Error

		if err != nil {
			log.Panicln("Error getting user and updating", err)
		}
		c.JSON(200, gin.H{
			"email":         foundUser.Email,
			"access_token":  foundUser.AccessToken,
			"refresh_token": foundUser.RefreshToken,
		})
		return

	}
}

func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		var passedUser models.User
		if err := c.BindJSON(&passedUser); err != nil {
			c.JSON(400, gin.H{"error": "Invalid parameters pass"})
			return
		}
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)
		user.Update(&passedUser, &*db)
		c.JSON(200, gin.H{"message": "Successfully Updated"})
		return
	}
}

func UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)
		serializedUser := user.UserDetailSerializer()
		fmt.Println(*serializedUser.AccessToken)
		fmt.Println(*serializedUser.Password)
		data, _ := json.MarshalIndent(serializedUser, "", "")
		c.JSON(200, gin.H{"message": "User Detail", "data": string(data)})
		return
	}
}
