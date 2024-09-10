package models

import (
	// "quickstart/db"
	"time"

	"gorm.io/gorm"
)

// User โมเดลสำหรับตาราง users
type User struct {
	UserID    uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string         `gorm:"size:100;not null;unique" json:"username"`
	Email     string         `gorm:"size:100;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Tokens    []Token        `gorm:"foreignKey:UserID" json:"tokens"` // เชื่อมโยงกับตาราง tokens
	Portfolio []Portfolio    `gorm:"foreignKey:UserID"`               // ความสัมพันธ์กับ Portfolio
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// func MigrateUser() error {
// 	return db.DB.AutoMigrate(&User{})
// }
