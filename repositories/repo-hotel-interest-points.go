package repositories

import "gorm.io/gorm"

// แท็ก
type HotelInterestPoints struct {
	gorm.Model
	HotelID           uint
	FacilityCode      int    `json:"facilityCode"`
	FacilityGroupCode int    `json:"facilityGroupCode"`
	Order             int    `json:"order"`
	PoiName           string `json:"poiName"`
	Distance          string `json:"distance"`
}
