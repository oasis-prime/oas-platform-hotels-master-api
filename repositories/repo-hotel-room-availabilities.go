package repositories

import "gorm.io/gorm"

// วันเวลาเปิดขายห้องพัก
type HotelRoomAvailabilities struct {
	gorm.Model
	HotelID uint
}
