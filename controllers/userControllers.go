package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"snapchat-clone/models"
	"snapchat-clone/serializers"
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
		serializedData := serializers.LoginSerializer(&foundUser)
		c.JSON(200, serializedData)
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
		serializedData := serializers.UserDetailSerializer(&user)
		c.JSON(200, serializedData)
		return
	}
}

func ProfileDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// Check if the user exists
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)

		var profile models.Profile

		// find the user profile
		err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Find(&profile).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		serializedData := serializers.ProfileDetailSerializer(&profile, &user)
		c.JSON(200, serializedData)
		return
	}
}

func ProfileUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		// Check if the user exists
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)

		// profile passed from posted data
		var profile models.Profile
		if err := c.BindJSON(&profile); err != nil {
			c.JSON(400, gin.H{"error": "Invalid parameters pass"})
			return
		}
		if profile.ProfileImageURL != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("profile_image_url", profile.ProfileImageURL).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.BackgroundImageURL != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).
				Update("background_image_url", profile.BackgroundImageURL).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.GhostMode != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("ghost_mode", profile.GhostMode).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.SeeLocation != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("see_location", profile.SeeLocation).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.LocationExceptFriends != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("location_except_friends", profile.LocationExceptFriends).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.ALlFriends != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("location_all_friends", profile.ALlFriends).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		if profile.TwoFactor != nil {
			err := db.Model(&profile).Where(&models.Profile{UserID: user.ID}).Update("two_factor", profile.TwoFactor).Error
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(200, gin.H{"messages": "Successfully updated profile"})
		return
	}
}
