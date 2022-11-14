package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Conversation struct {
	/* Once a user is connected to the websocket we add the user to this conversation*/
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name"`
	OnlineUSERS []User    `json:"online_users" gorm:"many2many:online_users"`
	Timestamp   time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate Setting the uuid before creating this stuff
func (u *Conversation) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type Message struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	FromUserID User      `json:"from_user"  gorm:"foreignKey:ID;"`
	ToUserID   User      `json:"to_user"  gorm:"foreignKey:ID;"`
	Content    string    `json:"content" validate:"max=2000,min=1,required"`
	Read       bool      `json:"read" gorm:"type:bool;default:false;"`
	Timestamp  time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate Setting the uuid before creating this stuff
func (u *Message) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
