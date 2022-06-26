package repositories

import "gorm.io/gorm"

// ราคาห้องพัก
type HotelRoomRates struct {
	gorm.Model
	HotelID uint
}
