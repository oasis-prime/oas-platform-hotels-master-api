package repositories

import (
	"gorm.io/gorm"
)

// ห้องพัก
type HotelRoomStayFacilities struct {
	gorm.Model
	HotelID           uint
	RoomStayID        uint
	FacilityCode      int `json:"facilityCode"`
	FacilityGroupCode int `json:"facilityGroupCode"`
	Number            int `json:"number"`
}
