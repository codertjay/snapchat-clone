package stories

import (
	"github.com/google/uuid"
	"snapchat-clone/users"
	"time"
)

type Story struct {
	ID         uuid.UUID    `json:"id" gorm:"primaryKey;"`
	User       users.User   `json:"user" gorm:"foreignKey:ID;"`
	File       string       `json:"file"`
	Name       string       `json:"name" gorm:"type:string; size:250;"`
	ViewCounts []users.User `json:"view_counts" gorm:"many2many:view_counts;"`
	Timestamp  time.Time    `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
