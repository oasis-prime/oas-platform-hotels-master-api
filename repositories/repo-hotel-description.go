package repositories

import (
	"gorm.io/gorm"
)

// รายละเอียดโรงแรม
type HotelDescription struct {
	gorm.Model
	HotelID  uint
	Language string
	Value    string
}