package users

import (
	"github.com/google/uuid"
	"time"
)

// User /* The user models struct */
type User struct {
	ID           uuid.UUID  `json:"id" gorm:"primaryKey;"`
	Name         *string    `json:"name" validate:"required,max=250,min=2" gorm:"size:250;"`
	Email        *string    `json:"email" validate:"required,max=100,min=5" gorm:"type:varchar(100);unique_index;"`
	Phone        *string    `json:"phone" validate:"required,max=20,min=3" gorm:"type:varchar(20);unique_index;"`
	Birthday     *time.Time `json:"birthday"`
	CreatedAt    *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP;"`
	Timestamp    *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
	Password     *string    `json:"password" validate:"required,max=250,min=5"`
	AccessToken  *string    `json:"access_token"`
	RefreshToken *string    `json:"refresh_token"`
}

// Profile /*  Profile*/
type Profile struct {
	//	the one to one relationship
	ID                    uuid.UUID  `json:"id" gorm:"primaryKey"`
	User                  *string    `json:"user" gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	UserID                uuid.UUID  `json:"user_id" gorm:"not null;"`
	ProfileImage          *string    `json:"profile_image"`
	BackgroundImage       *string    `json:"background_image"`
	GhostMode             bool       `json:"ghost_mode" gorm:"default:false;type:bool;"`
	SeeLocation           *string    `json:"see_location" validate:"omitempty,eq=FRIENDS|eq=EXCEPT_FRIENDS" gorm:"default:FRIENDS;"`
	LocationALlFriends    []User     `json:"location_all_friends" gorm:"many2many:location_all_friends;"`
	LocationExceptFriends []User     `json:"location_except_friends" gorm:"many2many:location_except_friends;"`
	TwoFactor             bool       `json:"two_factor" gorm:"default:false;type:bool;"`
	Timestamp             *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
