package repositories

import (
	"gorm.io/gorm"
)

// ชื่อโรงแรม
type HotelName struct {
	gorm.Model
	HotelID  uint
	Language string
	Value    string
}
