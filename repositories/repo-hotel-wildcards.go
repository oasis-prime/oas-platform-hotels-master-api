package repositories

import "gorm.io/gorm"

// แท็ก
type HotelWildcards struct {
	gorm.Model
	HotelID              uint
	RoomType             string `json:"roomType"`
	RoomCode             string `json:"roomCode"`
	CharacteristicCode   string `json:"characteristicCode"`
	HotelRoomDescription struct {
		Content string `json:"content"`
	} `json:"hotelRoomDescription"`
}
