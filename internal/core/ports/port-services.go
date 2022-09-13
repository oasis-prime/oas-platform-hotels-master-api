package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
)

type MemberService interface {
	PublisherVerifyEmail() (err error)
}

type HotelbedsService interface {
	AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error)
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
}

type GooglePlaceService interface {
	Queryautocomplete(condition *googledm.QueryautocompleteRequest) (res *googledm.QueryautocompleteResponse, err error)
}
