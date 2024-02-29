package entity

import (
	"time"
)

type Answer struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    int       `json:"user_id" gorm:""`
	Answers   string    `gorm:"column:answers"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
