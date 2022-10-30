package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Spotlight struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;"`
	User       User      `json:"user" gorm:"foreignKey:ID"`
	File       string    `json:"file"`
	Name       string    `json:"name" gorm:"type:string; size:250;"`
	ViewCounts []User    `json:"view_counts" gorm:"many2many:view_counts"`
	Timestamp  time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}

// BeforeCreate Setting the uuid before creating this stuff
func (u *Spotlight) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
