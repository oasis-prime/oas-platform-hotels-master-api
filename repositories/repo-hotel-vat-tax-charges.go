package repositories

import "gorm.io/gorm"

// ค่าธรรมเนียม
type HotelsVatTaxCharges struct {
	gorm.Model
	HotelID     uint
	Name        []*HotelName        `gorm:"foreignKey:HotelID"`
	Description []*HotelDescription `gorm:"foreignKey:HotelID"`
	IsActive    bool
}
