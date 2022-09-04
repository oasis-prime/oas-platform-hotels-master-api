package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
)

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*customerrepo.CustomerMember, totalRows int64, err error)
}

type HotelsRepository interface {
	GetAll(condition hoteldm.GetAllRequestBasic) (results []*hotelrepo.Hotels, totalRows int64, err error)
	Get(argID uint) (record *hotelrepo.Hotels, err error)
	CreateBatch(record *[]hotelrepo.Hotels) (RowsAffected int64, err error)
	Create(record *hotelrepo.Hotels) (result *hotelrepo.Hotels, RowsAffected int64, err error)
	Update(argID uint, updated *hotelrepo.Hotels) (result *hotelrepo.Hotels, RowsAffected int64, err error)
	Delete(argID uint32) (rowsAffected int64, err error)
	GetByCoordinates(condition hoteldm.GetByCoordinatesRequest) (record []*hotelrepo.Hotels, totalRows int64, err error)
}

type HotelNameRepository interface {
	GetAll(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error)
}

type HotelImagesRepository interface {
	GetAll(condition hoteldm.GetAllHotelImagesRequest) (results []*hotelrepo.HotelImages, totalRows int64, err error)
}

type HotelDescriptionRepository interface {
	GetAll(condition hoteldm.GetAllHotelDescriptionRequest) (results []*hotelrepo.HotelDescription, totalRows int64, err error)
}

type HotelInterestPointsRepository interface {
	GetAll(condition hoteldm.GetAllHotelInterestPointsRequest) (results []*hotelrepo.HotelInterestPoints, totalRows int64, err error)
}

type HotelIssuesRepository interface {
	GetAll(condition hoteldm.GetAllHotelIssuesRequest) (results []*hotelrepo.HotelIssues, totalRows int64, err error)
}

type HotelFacilityRepository interface {
	GetAll(condition hoteldm.GetAllHotelFacilityRequest) (results []*hotelrepo.HotelFacility, totalRows int64, err error)
}

type HotelRoomsRepository interface {
	GetAll(condition hoteldm.GetAllHotelRoomsRequest) (results []*hotelrepo.HotelRooms, totalRows int64, err error)
}

type HotelPhonesRepository interface {
	GetAll(condition hoteldm.GetAllHotelPhoneRequest) (results []*hotelrepo.HotelPhone, totalRows int64, err error)
}

type HotelCityRepository interface {
	GetAll(condition hoteldm.GetAllHotelCityRequest) (results []*hotelrepo.HotelCity, totalRows int64, err error)
}

type HotelAddressRepository interface {
	GetAll(condition hoteldm.GetAllHotelAddressRequest) (results []*hotelrepo.HotelAddress, totalRows int64, err error)
}

type HotelbedsBookingRepository interface {
}
