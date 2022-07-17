package repositories

import "gorm.io/gorm"

// รูปภาพ
type HotelImages struct {
	gorm.Model
	HotelID            uint
	ImageTypeCode      string `json:"imageTypeCode"`
	Path               string `json:"path"`
	Order              int    `json:"order"`
	VisualOrder        int    `json:"visualOrder"`
	RoomCode           string `json:"roomCode,omitempty"`
	RoomType           string `json:"roomType,omitempty"`
	CharacteristicCode string `json:"characteristicCode,omitempty"`
}
