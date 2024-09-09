package models

import (
    "time"
    "gorm.io/gorm"
)

// User โมเดลสำหรับตาราง users
type User struct {
    UserID    uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
    Username  string         `gorm:"size:100;not null;unique" json:"username"`
    Email     string         `gorm:"size:100;not null;unique" json:"email"`
    Password  string         `gorm:"size:255;not null" json:"password"`
    CreatedAt time.Time      `json:"created_at"`  // จะถูกตั้งค่าอัตโนมัติเมื่อสร้าง
    UpdatedAt time.Time      `json:"updated_at"`  // จะถูกอัปเดตอัตโนมัติเมื่อแก้ไข
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
