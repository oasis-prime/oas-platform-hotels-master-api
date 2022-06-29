package repositories

import "gorm.io/gorm"

// แท็ก
type HotelTags struct {
	gorm.Model
	HotelID uint
	Name    string
}
