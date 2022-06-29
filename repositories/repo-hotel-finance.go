package repositories

import (
	"gorm.io/gorm"
)

// การเงิน
type HotelFinance struct {
	gorm.Model
	HotelID        uint
	CommissionRate float64
}
