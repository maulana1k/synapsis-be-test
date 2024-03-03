package models

import "gorm.io/gorm"

type Cart struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID string `json:"userId"`
	User   User   `gorm:"foreignKey:UserID"`
}

type CartItems struct {
	gorm.Model
	CartID    uint    `json:"cartId"`
	Cart      Cart    `gorm:"foreignKey:CartID"`
	ProductID uint    `json:"productId"`
	Products  Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `json:"quantity"`
}
