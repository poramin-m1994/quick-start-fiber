package models

import (
	"time"

	"gorm.io/gorm"
)

// Portfolio บันทึกการซื้อหุ้นของผู้ใช้
type Portfolio struct {
	PortfolioID     uint              `gorm:"primaryKey;autoIncrement" json:"portfolio_id"`
	UserID          uint              `gorm:"not null" json:"user_id"`
	Portfolioname   string            `gorm:"size:100;not null;unique" json:"portfolio_name"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"-"`
	PortfolioDetail []PortfolioDetail `gorm:"foreignKey:PortfolioDetailID"` // ความสัมพันธ์กับประวัติการซื้อ
}
