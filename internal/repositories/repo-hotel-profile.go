package repositories

import "gorm.io/gorm"

// ตั้งค่า
type HotelProfile struct {
	gorm.Model
	HotelID uint
	Name    string
}
