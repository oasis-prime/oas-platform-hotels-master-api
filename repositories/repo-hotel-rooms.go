package repositories

import (
	"gorm.io/gorm"
)

// ห้องพัก
type HotelRooms struct {
	gorm.Model
	HotelID            uint
	RoomCode           string `json:"roomCode"`
	IsParentRoom       bool   `json:"isParentRoom"`
	MinPax             int    `json:"minPax"`
	MaxPax             int    `json:"maxPax"`
	MaxAdults          int    `json:"maxAdults"`
	MaxChildren        int    `json:"maxChildren"`
	MinAdults          int    `json:"minAdults"`
	RoomType           string `json:"roomType"`
	CharacteristicCode string `json:"characteristicCode"`
}
