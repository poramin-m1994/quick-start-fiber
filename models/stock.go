package models

import (
	"time"

	"gorm.io/gorm"
)

// Stock ตารางสำหรับข้อมูลหุ้น
type Stock struct {
	StockID         uint              `gorm:"primaryKey;autoIncrement" json:"stock_id"`
	Symbol          string            `gorm:"size:10;not null;unique" json:"symbol"` // สัญลักษณ์หุ้น (เช่น AMGN, MSFT)
	Name            string            `gorm:"size:255;not null" json:"name"`         // ชื่อหุ้น (เช่น Amgen Inc)
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"-"`
	PortfolioDetail []PortfolioDetail `gorm:"foreignKey:StockID"` // ความสัมพันธ์กับ Portfolio
}
