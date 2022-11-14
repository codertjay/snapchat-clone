package serializers

import (
	"github.com/google/uuid"
	"snapchat-clone/models"
)

/* Note serializers with no function are used for post request while serializers with function are used for get request*/

// SendAndAcceptFriendRequestSerializer  /* This serializer is used for sending friend request*/
type SendAndAcceptFriendRequestSerializer struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

// FriendRequestListSerializer  /* This serializer is meant to list all the friend received by the user*/
func FriendRequestListSerializer(friendRequests *models.FriendRequest) models.FriendRequest {
	var fromUser *models.User
	var toUser *models.User
	if *&friendRequests.FromUser != nil {
		fromUser = MinimumUserDetailSerializer(*&friendRequests.FromUser)
	}

	if *&friendRequests.ToUser != nil {
		toUser = MinimumUserDetailSerializer(*&friendRequests.ToUser)
	}
	serialized := models.FriendRequest{
		ID:         friendRequests.ID,
		Accepted:   friendRequests.Accepted,
		FromUser:   fromUser,
		ToUser:     toUser,
		FromUserID: friendRequests.FromUserID,
		ToUserID:   friendRequests.ToUserID,
		Timestamp:  friendRequests.Timestamp,
	}
	return serialized
}
