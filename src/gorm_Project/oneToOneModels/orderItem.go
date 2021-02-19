package oneToOneModels

import (
	"github.com/jinzhu/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID  uint
	ItemID   uint
	Item     Item
	Quantity int
}
