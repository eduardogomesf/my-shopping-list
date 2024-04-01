package models

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingListModel struct {
	gorm.Model
	ID         string     `gorm:"type:uuid;primary_key;column:id"`
	Name       string     `gorm:"type:varchar(255);column:name"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	FinishedAt *time.Time `gorm:"column:finished_at"`
	IsFinished bool       `gorm:"default:false;column:is_finished"`
}
