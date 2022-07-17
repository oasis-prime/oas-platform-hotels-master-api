package repositories

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type HotelsFacility struct {
	gorm.Model
	FacilityCode      int    `json:"facilityCode"`
	FacilityGroupCode int    `json:"facilityGroupCode"`
	Order             int    `json:"order"`
	IndYesOrNo        bool   `json:"indYesOrNo,omitempty"`
	Number            int    `json:"number,omitempty"`
	Voucher           bool   `json:"voucher"`
	IndLogic          bool   `json:"indLogic,omitempty"`
	IndFee            bool   `json:"indFee,omitempty"`
	Distance          int    `json:"distance,omitempty"`
	TimeFrom          string `json:"timeFrom,omitempty"`
	TimeTo            string `json:"timeTo,omitempty"`
	DateTo            string `json:"dateTo,omitempty"`
}
