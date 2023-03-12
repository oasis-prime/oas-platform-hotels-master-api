package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/domain/customerdm"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-firebase-core/domain/firebasedm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
)

type MemberService interface {
	PublisherVerifyEmail(condition domain.PublisherVerifyEmail) (err error)
	MemberRegister(condition firebasedm.MemberRegister) (err error)
}

type PaymentService interface {
	Generate(condition *chillpaydm.PaylinkGenerateRequest) (response *chillpaydm.PaylinkGenerateResponse, err error)
	CreatePayment(record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, err error)
	GetPayment(payLinkId int) (result *customerrepo.CustomerPayment, err error)
	GetChillPay(argCode uint) (response *chillpaydm.PaylinkGenerateResponse, err error)
	BookingMail(condition *domain.PublisherBookingEmail) (err error)
	UpdatePayment(argCode uint, record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, err error)
	GetDetailByTransctionID(condition *chillpaydm.TransctionIDRequest) (response *chillpaydm.DetailByTransctionResponse, err error)
}

type HotelbedsService interface {
	AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error)
	CheckRate(condition *hotelbedsdm.CheckRatesRequest) (res *hotelbedsdm.CheckRatesResponse, err error)
	Booking(condition *hotelbedsdm.BookingsRequest) (res *hotelbedsdm.BookingsResponse, err error)
}

type HotelsService interface {
	GetHotels(input model.HotelsInput) (record []*hotelrepo.Hotels, totalRows int64, err error)
	GetHotel(input model.HotelInput) (record *hotelrepo.Hotels, err error)
	GetHotelName(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error)
	GetHotelImages(condition hoteldm.GetAllHotelImagesRequest) (results []*hotelrepo.HotelImages, totalRows int64, err error)
	GetHotelFacility(condition hoteldm.GetAllHotelFacilityRequest) (results []*hotelrepo.HotelFacility, totalRows int64, err error)
	GetHotelAddress(condition hoteldm.GetAllHotelAddressRequest) (results []*hotelrepo.HotelAddress, totalRows int64, err error)
	GetHotelDescription(condition hoteldm.GetAllHotelDescriptionRequest) (results []*hotelrepo.HotelDescription, totalRows int64, err error)
	GetHotelInterestPoints(condition hoteldm.GetAllHotelInterestPointsRequest) (results []*hotelrepo.HotelInterestPoints, totalRows int64, err error)
	GetHotelIssues(condition hoteldm.GetAllHotelIssuesRequest) (results []*hotelrepo.HotelIssues, totalRows int64, err error)
	GetHotelPhones(condition hoteldm.GetAllHotelPhoneRequest) (results []*hotelrepo.HotelPhone, totalRows int64, err error)
	GetRooms(condition hoteldm.GetAllHotelRoomsRequest) (results []*hotelrepo.HotelRooms, totalRows int64, err error)
	GetRoomsFacility(condition hoteldm.GetAllHotelRoomFacilitiesRequest) (results []*hotelrepo.HotelRoomFacilities, totalRows int64, err error)
	GetRoomsStay(condition hoteldm.GetAllHotelRoomStayRequest) (results []*hotelrepo.HotelRoomStay, totalRows int64, err error)
	GetCity(condition hoteldm.GetAllHotelCityRequest) (results []*hotelrepo.HotelCity, totalRows int64, err error)
	GetCoordinates(condition hoteldm.GetAllHotelCoordinatesRequest) (results []*hotelrepo.HotelCoordinates, totalRows int64, err error)
}

type GooglePlaceService interface {
	Queryautocomplete(condition *googledm.QueryautocompleteRequest) (res *googledm.QueryautocompleteResponse, err error)
}

type BookingService interface {
	GetAll(condition customerdm.GetAllRequestBasic) (results []*customerrepo.CustomerBooking, totalRows int64, err error)
	Get(argID uint) (record *customerrepo.CustomerBooking, err error)
	Create(record *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error)
	Update(argID uint, updated *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error)
	GetPayLinkId(payLinkId string) (record *customerrepo.CustomerBooking, err error)
}

type MasterService interface {
	GetCities(input masterdm.GetAllRequestBasic) (results []*masterrepo.City, totalRows int64, err error)
}

type PopularService interface {
	GetAll(condition masterdm.GetCustomerPopularRequest) (results []*masterrepo.Popular, totalRows int64, err error)
	Get(condition masterdm.GetCustomerPopularRequest) (record *masterrepo.Popular, err error)
}

type TickerService interface {
	TickerGetAll(condition masterdm.GetAllRequestBasic) (results []*masterrepo.Ticker, totalRows int64, err error)
	TickerGet(argID uint) (record *masterrepo.Ticker, err error)
	CustomerTickerGetAll(condition customerdm.GetAllRequestBasic) (results []*customerrepo.CustomerTicker, totalRows int64, err error)
	CustomerTickerGet(argID uint) (record *customerrepo.CustomerTicker, err error)
}

type SubscriptionService interface {
}
