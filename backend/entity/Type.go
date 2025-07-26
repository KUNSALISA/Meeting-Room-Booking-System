package entity

import (
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	TypeName string

	Rooms []Room `gorm:"foreignKey:TypeID"`
}