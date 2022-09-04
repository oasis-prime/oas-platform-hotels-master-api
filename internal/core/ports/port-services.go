package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
)

type MemberService interface {
}

type HotelbedsService interface {
	AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error)
}

type HotelsService interface {
	GetHotel(input model.HotelsInput) (record []*hotelrepo.Hotels, totalRows int64, err error)
	GetHotelName(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error)
	GetHotelImages(condition hoteldm.GetAllHotelImagesRequest) (results []*hotelrepo.HotelImages, totalRows int64, err error)
	GetHotelFacility(condition hoteldm.GetAllHotelFacilityRequest) (results []*hotelrepo.HotelFacility, totalRows int64, err error)
}

type GooglePlaceService interface {
	Queryautocomplete(condition *googledm.QueryautocompleteRequest) (res *googledm.QueryautocompleteResponse, err error)
}
