package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Title       string    `gorm:"not null" json:"title"` // หัวข้อที่จะประชุม
	Description string    // อธิบายรายละเอียดการประชุมนัดหมาย
	Date        time.Time `gorm:"type:date"`
	StartTime   time.Time `gorm:"type:time"`
	EndTime     time.Time `gorm:"type:time"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	RoomID uint
	Room   Room `gorm:"foreignKey:RoomID"`

}
