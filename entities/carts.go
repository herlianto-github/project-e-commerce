package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID             uint
	Total_Product  int
	Total_price    int
	User_id        uint
	Detail_cart_ID []Detail_cart
}

type Detail_cart struct {
	gorm.Model
	ID           uint
	CartID       uint `gorm:"not unique"`
	ProductID    uint `gorm:"not unique"`
	Qty          int
	Price        int
	TotalPrice   int
	DateCheckout *time.Time
}
