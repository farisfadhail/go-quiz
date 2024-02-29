package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:100;column:username"`
	Email     string    `gorm:"size:255;column:email;uniqueIndex"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role;default:respondent"` // respondent, surveyCreator & admin
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Answers   []Answer  `json:"answers"`
}
