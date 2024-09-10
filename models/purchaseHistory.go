package models

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseHistory struct {
	PurchaseID       uint           `gorm:"primaryKey;autoIncrement" json:"purchase_id"`
	PortfolioID      uint           `gorm:"not null" json:"portfolio_id"`
	PurchaseDate     time.Time      `gorm:"not null" json:"purchase_date"`
	PurchasePrice    float64        `gorm:"not null" json:"purchase_price"`
	PurchaseQuantity float64        `gorm:"not null" json:"purchase_quantity"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
