package oneToOneModels

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Status     string
	OrderItems []OrderItem `gorm:"one2many:order_item"`
}
