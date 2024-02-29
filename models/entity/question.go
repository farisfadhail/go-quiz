package entity

import (
	"github.com/lib/pq"
	"time"
)

type Question struct {
	ID        uint           `gorm:"primaryKey"`
	Question  string         `gorm:"column:question"`
	Points    pq.StringArray `gorm:"type:text;column:points"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
