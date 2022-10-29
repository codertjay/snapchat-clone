package models

import (
	"github.com/google/uuid"
	"time"
)

type Story struct {
	ID         uuid.UUID  `json:"id" gorm:"primaryKey;"`
	User       *User      `json:"user" gorm:"constraint:OnDelete:CASCADE;"`
	UserID     *uuid.UUID `json:"user_id" gorm:"not null;"`
	File       *string    `json:"file"`
	Name       *string    `json:"name" gorm:"type:string; size:250;"`
	ViewCounts *[]User    `json:"view_counts" gorm:"many2many:view_counts;"`
	Timestamp  *time.Time `json:"timestamp" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;"`
}
