package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"snapchat-clone/models"
	"snapchat-clone/serializers"
	snapchat_clone "snapchat-clone/snapchat-clone/database"
	"time"
)

func SendFriendRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//  initialize the db connection
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		var friendRequest models.FriendRequest
		/* Convert json and binds it for golang to understand*/
		if err := c.BindJSON(&friendRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//	 logged in user
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)
		friendRequest.FromUserID = user.ID
		// initialize the not accepted since we are using pointer
		not_accepted := false
		friendRequest.Accepted = &not_accepted
		err := db.Table("friend_requests").Create(&friendRequest).Error
		if err != nil {
			c.JSON(500, gin.H{"error": "Error saving friend requests"})
			return
		}
		c.JSON(200, gin.H{"Message": "Successfully sent friend request"})
		return

	}
}

// ReceivedFriendRequests this returns all the friend request received by the user both accepted
// and not accepted, but it can also be queried with params
func ReceivedFriendRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//  initialize the db connection
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		//	 logged in user
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)
		// return all
		result := db.Table("friend_requests").Find("to_user_id", user.ID)
		c.JSON(200, result)
	}
}

func AcceptFriendRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		//	 initialize the db connection
		db := snapchat_clone.DBConnection()
		defer snapchat_clone.CloseDB()
		//	 logged in user
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)

		// initialize accept request serializer
		var accept_request_serializer serializers.AcceptRequestSerializer
		/* Convert json and binds it for golang to understand*/
		if err := c.BindJSON(&accept_request_serializer); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// find the user we are adding to friends
		var other_user models.User
		err := db.Table("users").Find("id", accept_request_serializer.ID).First(&other_user).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "Other user not found"})
			return

		}
		err = db.Table("friend_requests").Find("from_user_id", accept_request_serializer.ID).
			Update("accepted", true).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "An error occurred acceptation user request"})
			return
		}

		err = db.Table("profiles").Find("id", user.ID).
			Association("LocationALlFriends").
			Append(&other_user)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error appending to user"})
			return
		}
		c.JSON(400, gin.H{"message": "Successfully accepted friend requests"})
		return

	}
}
