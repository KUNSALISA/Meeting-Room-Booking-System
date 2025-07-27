package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CodeName    string `gorm:"unique"`
	Password    string
	Firstname   string
	Lastname    string
	Image       string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`

	RoleID uint
	Role   Role `gorm:"foreignKey:RoleID"`

	Bookings []Booking `gorm:"foreignKey:UserID"`
}

