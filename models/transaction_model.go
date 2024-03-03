package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	// UserID string  `json:"user_id"`
	// User   User    `gorm:"foreignKey:UserID"`
	CartID uint    `json:"cartId"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
