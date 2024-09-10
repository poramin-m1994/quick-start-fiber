package models

import (
	"time"

	"gorm.io/gorm"
)

// Portfolio บันทึกการซื้อหุ้นของผู้ใช้
type PortfolioDetail struct {
	PortfolioDetailID uint              `gorm:"primaryKey;autoIncrement" json:"portfolio_id"`
	StockID           uint              `gorm:"not null" json:"stock_id"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
	DeletedAt         gorm.DeletedAt    `gorm:"index" json:"-"`
	PurchaseHistories []PurchaseHistory `gorm:"foreignKey:PortfolioID"` // ความสัมพันธ์กับประวัติการซื้อ
}
