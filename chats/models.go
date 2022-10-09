package chats

import (
	"github.com/google/uuid"
	"snapchat-clone/users"
	"time"
)

type Conversation struct {
	/* Once a user is connected to the websocket we add the user to this conversation*/

	ID          uuid.UUID    `json:"id" gorm:"type:uuid;primary_key;"`
	Name        string       `json:"name"`
	OnlineUSERS []users.User `json:"online_users" gorm:"many2many:online_users"`
	Timestamp   time.Time    `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

type Message struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;"`
	FromUserID users.User `json:"from_user"  gorm:"foreignKey:ID;"`
	ToUserID   users.User `json:"to_user"  gorm:"foreignKey:ID;"`
	Content    string     `json:"content" validate:"max=2000,min=1,required"`
	Read       bool       `json:"read" gorm:"type:bool;default:false;"`
	Timestamp  time.Time  `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

type FriendRequest struct {
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	FromUserID users.User `json:"from_user" gorm:"foreignKey:ID;"`
	ToUserID   users.User `json:"to_user" gorm:"foreignKey:ID;"`
	Accepted   bool       `json:"accepted" gorm:"type:bool;default:false;"`
	Timestamp  time.Time  `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
