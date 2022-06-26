package repositories

import (
	"gorm.io/gorm"
)

// ห้องพัก
type HotelRooms struct {
	gorm.Model
	HotelID  uint
	Language string
	Value    string
}
