package repositories

import (
	"gorm.io/gorm"
)

// ห้องพัก
type HotelRoomStay struct {
	gorm.Model
	HotelID     uint
	RoomID      uint
	StayType    string `json:"stayType"`
	Order       string `json:"order"`
	Description string `json:"description"`
}
