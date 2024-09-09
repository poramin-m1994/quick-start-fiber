package models

import (
	// "quickstart/db"
	"time"

	"gorm.io/gorm"
)

// Token โมเดลสำหรับตาราง token
type Token struct {
    TokenID   uint           `gorm:"primaryKey;autoIncrement" json:"token_id"`
    UserID    uint           `gorm:"not null" json:"user_id"`
    Token     string         `gorm:"size:255;not null" json:"token"`
    ExpiresAt time.Time      `json:"expires_at"`
    CreatedAt time.Time      `json:"created_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// func MigrateToken() error {
// 	return db.DB.AutoMigrate(&Token{})
// }
