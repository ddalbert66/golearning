package projectModels

import (
	"github.com/jinzhu/gorm"
)

type EnumType struct {
	gorm.Model
	Code       string
	Name       string
	CreateUser string
	UpdateUser string
	Version    uint
	Enums      []Enumeration `gorm:"foreignKey:EnumTypeID"`
}
