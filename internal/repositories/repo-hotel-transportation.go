package repositories

import "gorm.io/gorm"

// วิธีเดินทาง
type HotelTransportation struct {
	gorm.Model
	HotelID uint
	Name    string
}
