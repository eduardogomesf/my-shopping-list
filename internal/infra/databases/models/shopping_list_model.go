package models

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingListModel struct {
	gorm.Model
	ID         string `gorm:"type:uuid;primary_key"`
	Name       string `gorm:"type:size:255"`
	CreatedAt  time.Time
	FinishedAt *time.Time
	IsFinished bool `gorm:"default:false"`
}
