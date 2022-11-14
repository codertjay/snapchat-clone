package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AllFriend Friend /* This contains list of all the friends a user have */
// fixme: try to create a friend schema that enable both users to use
type AllFriend struct {
	ID          uuid.UUID  `json:"id"  gorm:"primaryKey;"`
	User        *User      `json:"user"  gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	Friend      *User      `json:"friend"  gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	SeeLocation *bool      `json:"seeLocation" gorm:"default:false;type:bool;"`
	Timestamp   *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate Setting the uuid before creating this stuff
func (u *AllFriend) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

// FriendRequest /* this contains list of friend request sent and received .
// so through this part the user is able to */
type FriendRequest struct {
	ID         uuid.UUID  `json:"id"  gorm:"primaryKey;"`
	Accepted   *bool      `json:"accepted" gorm:"default:false;type:bool;"`
	FromUser   *User      `json:"from_user"   gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	ToUser     *User      `json:"to_user"  gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	FromUserID uuid.UUID  `json:"from_user_id" gorm:"not null;"`
	ToUserID   uuid.UUID  `json:"to_user_id" validate:"required" gorm:"not null;"`
	Timestamp  *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate Setting the uuid before creating this stuff
func (u *FriendRequest) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
