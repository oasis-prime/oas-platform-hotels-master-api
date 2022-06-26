package repositories

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type Hotels struct {
	gorm.Model
	Name        []*HotelName        `gorm:"foreignKey:HotelID"`
	Description []*HotelDescription `gorm:"foreignKey:HotelID"`
	Status      bool
}
