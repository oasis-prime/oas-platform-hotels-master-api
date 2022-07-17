package repositories

import (
	"gorm.io/gorm"
)

// ชื่อโรงแรม
type HotelName struct {
	gorm.Model
	HotelID uint
	Key     string
	Value   string
}
