package models

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	/* Once a user is connected to the websocket we add the user to this conversation*/

	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name"`
	OnlineUSERS []User    `json:"online_users" gorm:"many2many:online_users"`
	Timestamp   time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	FromUserID User      `json:"from_user"  gorm:"foreignKey:ID;"`
	ToUserID   User      `json:"to_user"  gorm:"foreignKey:ID;"`
	Content    string    `json:"content" validate:"max=2000,min=1,required"`
	Read       bool      `json:"read" gorm:"type:bool;default:false;"`
	Timestamp  time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// FriendRequest /* this contains list of friend request sent and received .
// so through this part the user is able to */
type FriendRequest struct {
	ID         uuid.UUID  `json:"id"  gorm:"primaryKey"`
	Accepted   *bool      `json:"accepted" gorm:"default:false;type:bool;"`
	FromUser   *User      `json:"from_user"   gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	ToUser     *User      `json:"to_user"  gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	FromUserID uuid.UUID  `json:"from_user_id" gorm:"not null;"`
	ToUserID   uuid.UUID  `json:"to_user_id" validate:"required" gorm:"not null;"`
	Timestamp  *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
