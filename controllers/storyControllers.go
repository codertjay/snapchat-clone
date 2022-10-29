package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"snapchat-clone/models"
	snapchat_clone "snapchat-clone/snapchat-clone/database"
	"time"
)

func UserStoryList() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**/
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Used to list the story's that are available for the user
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()

		// the authenticated user
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)

		//
		// setting the profile what we initialized
		var friends []models.User
		var stories []models.Story
		err := db.Table("profiles").
			Find("id", user.ID).
			Association("ALlFriends").Find(&friends).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "An error occurred getting user profile"})
			return
		}
		// now we have all the friends in a list so let's search and get the story's
		for _, item := range friends {
			var story models.Story
			err := db.Table("story").Where(&models.Story{UserID: &item.ID}).First(&story).Error
			if err != nil {
				c.JSON(400, gin.H{"error": "An error occurred getting story"})
				return
			}
			stories = append(stories, story)
		}
		c.JSON(200, stories)
		return

	}
}
