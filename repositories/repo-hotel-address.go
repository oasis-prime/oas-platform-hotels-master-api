package repositories

import (
	"gorm.io/gorm"
)

// ที่อยู่โรงแรม
type HotelAddress struct {
	gorm.Model
	HotelID  uint
	Language string
	Value    string
}
