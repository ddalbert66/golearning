package oneToOneModels

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ItemName string
	Amount   float32
}
