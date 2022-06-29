package repositories

import "gorm.io/gorm"

// เข้าร่วมแคมเปญ
type HotelCampaign struct {
	gorm.Model
	HotelID uint
}
