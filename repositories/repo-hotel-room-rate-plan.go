package repositories

import "gorm.io/gorm"

// ราคาห้องพัก
type HotelRoomRatePlan struct {
	gorm.Model
	HotelID uint
}
