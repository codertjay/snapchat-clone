package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	snapchat_clone "snapchat-clone/snapchat-clone/database"
	"time"
)

// Note : Fields with omit empty are not allowed to be shown in the frontend

// User /* The user models struct */
type User struct {
	ID           uuid.UUID  `json:"id" gorm:"primaryKey;"`
	Name         *string    `json:"name" validate:"required,max=250,min=2" gorm:"size:250;"`
	Email        *string    `json:"email" validate:"required,max=100,min=5" gorm:"type:varchar(100);unique_index;"`
	Phone        *string    `json:"phone" validate:"required,max=20,min=3" gorm:"type:varchar(20);unique_index;"`
	Birthday     *time.Time `json:"birthday"`
	CreatedAt    *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP;"`
	Timestamp    *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
	Password     *string    `json:"password,omitempty" validate:"required,max=250,min=5"`
	AccessToken  *string    `json:"access_token,omitempty"`
	RefreshToken *string    `json:"refresh_token,omitempty"`
}

func (u *User) UserDetailSerializer() *User {
	var user = User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
		Timestamp: u.Timestamp,
	}
	return &user
}

func (u *User) FindAll() User {
	var users User
	db := snapchat_clone.DBConnection()
	snapchat_clone.CloseDB()
	err := db.Find(&users).Error
	if err != nil {
		return users
	}
	return users
}

// Update /* This save the user data*/
func (u *User) Update(user *User, db *gorm.DB) {
	if user.Name != nil {
		err := db.Model(&user).Where(&User{ID: u.ID}).Update("name", user.Name).Error
		if err != nil {
			log.Panicln("Error occurred updating user", err)
			return
		}
	}
	if user.Email != nil {
		err := db.Model(&user).Where(&User{ID: u.ID}).Update("email", user.Email).Error
		if err != nil {
			log.Panicln("Error occurred", err)
			return
		}
	}
	if user.Phone != nil {
		err := db.Model(&user).Where(&User{ID: u.ID}).Update("phone", user.Phone).Error
		if err != nil {
			log.Panicln("Error occurred updating user", err)
			return
		}
	}
	if user.Birthday != nil {
		err := db.Model(&user).Where(&User{ID: u.ID}).Update("phone", user.Phone).Error
		if err != nil {
			log.Panicln("Error occurred updating user", err)
			return
		}
	}
}

// Profile /*  Profile*/
type Profile struct {
	//	the one to one relationship
	ID                    uuid.UUID  `json:"id" gorm:"primaryKey"`
	User                  *User    `json:"user" gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	UserID                uuid.UUID  `json:"user_id" gorm:"not null;"`
	ProfileImage          *string    `json:"profile_image"`
	BackgroundImage       *string    `json:"background_image"`
	GhostMode             *bool      `json:"ghost_mode" gorm:"default:false;type:bool;"`
	SeeLocation           *string    `json:"see_location" validate:"omitempty,eq=FRIENDS|eq=EXCEPT_FRIENDS" gorm:"default:FRIENDS;"`
	LocationALlFriends    []User     `json:"location_all_friends" gorm:"many2many:location_all_friends;"`
	LocationExceptFriends []User     `json:"location_except_friends" gorm:"many2many:location_except_friends;"`
	TwoFactor             bool       `json:"two_factor" gorm:"default:false;type:bool;"`
	Timestamp             *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
