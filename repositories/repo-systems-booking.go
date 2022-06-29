package repositories

import (
	"gorm.io/gorm"
)

// จองห้องพัก
type SystemsBooking struct {
	gorm.Model
	HotelID uint
}
