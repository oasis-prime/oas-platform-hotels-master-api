package repositories

import "gorm.io/gorm"

// แท็ก
type HotelTerminals struct {
	gorm.Model
	HotelID      uint
	TerminalCode string `json:"terminalCode"`
	Distance     int    `json:"distance"`
}
