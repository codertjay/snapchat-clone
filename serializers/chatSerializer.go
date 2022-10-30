package serializers

import "github.com/google/uuid"

type AcceptRequestSerializer struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Accepted bool      `json:"accepted"`
}

// FriendRequestSerializer /* This serializer is used for sending friend request*/
type FriendRequestSerializer struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
