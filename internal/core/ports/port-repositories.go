package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/customerdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-firebase-core/domain/firebasedm"
)

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*customerrepo.CustomerMember, totalRows int64, err error)
}

type PaymentRepository interface {
	GetAll(condition customerdm.GetAllPaymentRequest) (results []*customerrepo.CustomerPayment, totalRows int64, err error)
	Get(argCode uint) (record *customerrepo.CustomerPayment, err error)
	GetPayLinkId(payLinkId int) (record *customerrepo.CustomerPayment, err error)
	Create(record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, RowsAffected int64, err error)
	Update(argCode uint, updated *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, RowsAffected int64, err error)
}

type HotelsRepository interface {
	GetAll(condition hoteldm.GetAllHotelRequest) (results []*hotelrepo.Hotels, totalRows int64, err error)
	GetByCoordinates(condition hoteldm.GetByCoordinatesRequest) (record []*hotelrepo.Hotels, totalRows int64, err error)
	Get(condition hoteldm.GetHotelRequest) (record *hotelrepo.Hotels, err error)
}

type HotelNameRepository interface {
	GetAll(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error)
}

type HotelCoordinatesRepository interface {
	GetAll(condition hoteldm.GetAllHotelCoordinatesRequest) (results []*hotelrepo.HotelCoordinates, totalRows int64, err error)
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

type HotelRoomsFacilitiesRepository interface {
	GetAll(condition hoteldm.GetAllHotelRoomFacilitiesRequest) (results []*hotelrepo.HotelRoomFacilities, totalRows int64, err error)
}

type HotelRoomsStayRepository interface {
	GetAll(condition hoteldm.GetAllHotelRoomStayRequest) (results []*hotelrepo.HotelRoomStay, totalRows int64, err error)
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

type FirebaseMemberRepository interface {
	MemberRegister(condition firebasedm.MemberRegister) (result *firebasedm.MemberRegisterResponse, err error)
	MemberUpdate(argID string, condition firebasedm.MemberUpdate) (result *firebasedm.MemberRegisterResponse, err error)
}

type CustomerBookingRepository interface {
	GetAll(condition customerdm.GetAllRequestBasic) (results []*customerrepo.CustomerBooking, totalRows int64, err error)
	Get(argID uint) (record *customerrepo.CustomerBooking, err error)
	Create(record *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error)
	Update(argID uint, updated *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error)
	GetPayLinkId(payLinkId string) (record *customerrepo.CustomerBooking, err error)
}

type PopularRepository interface {
	GetAll(condition masterdm.GetCustomerPopularRequest) (results []*masterrepo.Popular, totalRows int64, err error)
	Get(condition masterdm.GetCustomerPopularRequest) (record *masterrepo.Popular, err error)
}
