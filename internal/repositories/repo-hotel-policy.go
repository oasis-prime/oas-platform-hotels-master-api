package repositories

import (
	"gorm.io/gorm"
)

// นโยบาย
type HotelPolicy struct {
	gorm.Model
	HotelID uint
}
