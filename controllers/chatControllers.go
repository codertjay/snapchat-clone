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
		/* I know you might think why don't I just use the model for the json stuff but. Somehow user might send field that could conflict with the
		existing field that i won't like to change from json but from my code*/
		var friendRequestSerializer serializers.FriendRequestSerializer
		/* Convert json and binds it for golang to understand*/
		if err := c.BindJSON(&friendRequestSerializer); err != nil {
			c.JSON(400, gin.H{"error": "Error converting to json"})
			return
		}
		// validating
		validatorErr := validate.Struct(friendRequestSerializer)
		if validatorErr != nil {
			c.JSON(400, gin.H{"error": validatorErr.Error()})
			return
		}
		//	 logged in user
		loggedInUser, _ := c.Get("user")
		user := loggedInUser.(models.User)
		// Check if the user is not sending friend request to him self
		if friendRequestSerializer.ID == user.ID {
			c.JSON(400, gin.H{"error": "You cant send friend request to yourself"})
			return
		}
		// initialize the model struct
		var friendRequest models.FriendRequest
		friendRequest.FromUserID = user.ID
		// using the id sent by the user for request
		friendRequest.ToUserID = friendRequestSerializer.ID
		// before creating i check if i have sent a friend request to the user before
		var count int64
		db.Table("friend_requests").
			Find(&models.FriendRequest{ToUserID: friendRequest.ToUserID, FromUserID: friendRequest.FromUserID}).
			Count(&count)
		if int(count) > 0 {
			c.JSON(400, gin.H{"error": "Already sent request to this user"})
			return
		}
		// check if the user is also my friend already
		db.Table("all_friends").
			Where(&models.Profile{UserID: user.ID}).
			Association("AllFriends").
			Find("user_id", friendRequest.ToUserID).Error()

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
		var acceptRequestSerializer serializers.AcceptRequestSerializer
		/* Convert json and binds it for golang to understand*/
		if err := c.BindJSON(&acceptRequestSerializer); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// find the user we are adding to friends
		var otherUser models.User
		err := db.Table("users").Find("id", acceptRequestSerializer.ID).First(&otherUser).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "Other user not found"})
			return

		}
		err = db.Table("friend_requests").Find("from_user_id", acceptRequestSerializer.ID).
			Update("accepted", true).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "An error occurred acceptation user request"})
			return
		}

		err = db.Table("profiles").Find("id", user.ID).
			Association("LocationALlFriends").
			Append(&otherUser)
		if err != nil {
			c.JSON(400, gin.H{"error": "Error appending to user"})
			return
		}
		c.JSON(400, gin.H{"message": "Successfully accepted friend requests"})
		return

	}
}
