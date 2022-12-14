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
	ID           uuid.UUID  `json:"id,omitempty" gorm:"primaryKey;type:uuid;"`
	Name         *string    `json:"name,omitempty" validate:"required,max=250,min=2" gorm:"size:250;"`
	Email        *string    `json:"email,omitempty" validate:"required,max=100,min=5" gorm:"type:varchar(100);unique_index;"`
	Phone        *string    `json:"phone,omitempty" validate:"required,max=20,min=3" gorm:"type:varchar(20);unique_index;"`
	Birthday     *time.Time `json:"birthday,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty" gorm:"default:CURRENT_TIMESTAMP;"`
	Timestamp    *time.Time `json:"timestamp,omitempty" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
	Password     *string    `json:"password,omitempty" validate:"required,max=250,min=5"`
	AccessToken  *string    `json:"access_token,omitempty"`
	RefreshToken *string    `json:"refresh_token,omitempty"`
}

// Setting the uuid before creating this stuff
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
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
	ID                    uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;"`
	User                  *User      `json:"user" gorm:"constraint:OnDelete:CASCADE;unique_index;"`
	UserID                uuid.UUID  `json:"user_id" gorm:"not null;"`
	ProfileImageURL       *string    `json:"profile_image_url"`
	BackgroundImageURL    *string    `json:"background_image_url"`
	GhostMode             *bool      `json:"ghost_mode" gorm:"default:false;type:bool;"`
	SeeLocation           *string    `json:"see_location" validate:"omitempty,eq=FRIENDS|eq=EXCEPT_FRIENDS" gorm:"default:FRIENDS;"`
	AllFriends            []User     `json:"all_friends" gorm:"many2many:all_friends;"`
	LocationExceptFriends []User     `json:"location_except_friends" gorm:"many2many:location_except_friends;"`
	TwoFactor             *bool      `json:"two_factor" gorm:"default:false;type:bool;"`
	Timestamp             *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate  setting the uuid of the value
func (u *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
