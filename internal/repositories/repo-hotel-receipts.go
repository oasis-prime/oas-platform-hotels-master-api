package repositories

import "gorm.io/gorm"

// ตั้งค่า
type HotelReceipts struct {
	gorm.Model
	HotelID uint
	Name    string
}
