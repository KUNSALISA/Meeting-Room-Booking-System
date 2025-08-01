package entity

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomName  string `gorm:"unique"`
	Location  string
	Capacity  uint
	Equipment pq.StringArray `gorm:"type:text[]"`

	TypeID uint
	Types   Type `gorm:"foreignKey:TypeID"`

	StatusID uint
	Status   Status `gorm:"foreignKey:StatusID"`

	Bookings []Booking `gorm:"foreignKey:RoomID"`
}
