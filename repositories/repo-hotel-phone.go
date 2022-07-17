package repositories

import "gorm.io/gorm"

// แท็ก
type HotelPhone struct {
	gorm.Model
	HotelID     uint
	PhoneNumber string `json:"phoneNumber"`
	PhoneType   string `json:"phoneType"`
}
