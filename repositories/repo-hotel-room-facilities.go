package repositories

import (
	"gorm.io/gorm"
)

// ห้องพัก
type HotelRoomFacilities struct {
	gorm.Model
	HotelID           uint
	RoomID            uint
	FacilityCode      int  `json:"facilityCode"`
	FacilityGroupCode int  `json:"facilityGroupCode"`
	IndLogic          bool `json:"indLogic"`
	Number            int  `json:"number"`
	Voucher           bool `json:"voucher"`
}
