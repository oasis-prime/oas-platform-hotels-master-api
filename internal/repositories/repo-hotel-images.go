package repositories

import "gorm.io/gorm"

// รูปภาพ
type HotelImages struct {
	gorm.Model
	HotelID uint
}
