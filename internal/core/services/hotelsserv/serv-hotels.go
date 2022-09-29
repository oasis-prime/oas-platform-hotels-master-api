package hotelsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/htenums"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	repoHotels               ports.HotelsRepository
	repoHotelName            ports.HotelNameRepository
	repoHotelDescription     ports.HotelDescriptionRepository
	repoHotelInterestPoints  ports.HotelInterestPointsRepository
	repoHotelIssues          ports.HotelIssuesRepository
	repoHotelFacility        ports.HotelFacilityRepository
	repoHotelRooms           ports.HotelRoomsRepository
	repoHotelRoomsFacilities ports.HotelRoomsFacilitiesRepository
	repoHotelRoomsStay       ports.HotelRoomsStayRepository
	repoHotelPhones          ports.HotelPhonesRepository
	repoHotelCity            ports.HotelCityRepository
	repoHotelAddress         ports.HotelAddressRepository
	repoHotelImages          ports.HotelImagesRepository
	repoCoordinates          ports.HotelCoordinatesRepository
}

func NewService(
	repoHotels ports.HotelsRepository,
	repoHotelName ports.HotelNameRepository,
	repoHotelDescription ports.HotelDescriptionRepository,
	repoHotelInterestPoints ports.HotelInterestPointsRepository,
	repoHotelIssues ports.HotelIssuesRepository,
	repoHotelFacility ports.HotelFacilityRepository,
	repoHotelRooms ports.HotelRoomsRepository,
	repoHotelRoomsFacilities ports.HotelRoomsFacilitiesRepository,
	repoHotelRoomsStay ports.HotelRoomsStayRepository,
	repoHotelPhones ports.HotelPhonesRepository,
	repoHotelCity ports.HotelCityRepository,
	repoHotelAddress ports.HotelAddressRepository,
	repoHotelImages ports.HotelImagesRepository,
	repoCoordinates ports.HotelCoordinatesRepository,
) *service {
	return &service{
		repoHotels:               repoHotels,
		repoHotelName:            repoHotelName,
		repoHotelDescription:     repoHotelDescription,
		repoHotelInterestPoints:  repoHotelInterestPoints,
		repoHotelIssues:          repoHotelIssues,
		repoHotelFacility:        repoHotelFacility,
		repoHotelRooms:           repoHotelRooms,
		repoHotelRoomsFacilities: repoHotelRoomsFacilities,
		repoHotelRoomsStay:       repoHotelRoomsStay,
		repoHotelPhones:          repoHotelPhones,
		repoHotelCity:            repoHotelCity,
		repoHotelAddress:         repoHotelAddress,
		repoHotelImages:          repoHotelImages,
		repoCoordinates:          repoCoordinates,
	}
}

func (svc *service) GetHotels(
	input model.HotelsInput,
) (record []*hotelrepo.Hotels, totalRows int64, err error) {
	pageSize := input.Pagination.PageSize
	if pageSize > 20 {
		pageSize = 20
	}

	if input.Geolocation != nil {
		return svc.repoHotels.GetByCoordinates(hoteldm.GetByCoordinatesRequest{
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				Page:     input.Pagination.Page,
				PageSize: pageSize,
				Language: (*langenums.Language)(&input.Language),
			},
			Latitude:  input.Geolocation.Latitude,
			Longitude: input.Geolocation.Longitude,
			Distance:  uint(input.Geolocation.Radius),
		})
	}

	if input.Keywords != nil {
		rooms, adults, children := 0, 0, 0

		if input.Occupancies != nil {
			rooms = input.Occupancies.Rooms
			adults = input.Occupancies.Adults
			children = input.Occupancies.Children
		}

		return svc.repoHotels.GetAll(hoteldm.GetAllHotelRequest{
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				Page:     input.Pagination.Page,
				PageSize: pageSize,
				Keyword:  &input.Keywords.Keyword,
				Language: (*langenums.Language)(&input.Language),
			},
			Rooms:    rooms,
			Adults:   adults,
			Children: children,
		})
	}

	return nil, 0, nil
}

func (svc *service) GetHotel(
	input model.HotelInput,
) (record *hotelrepo.Hotels, err error) {
	condition := hoteldm.GetHotelRequest{
		Code:     uint(input.Code),
		Type:     htenums.Hotelbeds,
		Language: langenums.Language(input.Language),
	}
	return svc.repoHotels.Get(condition)
}

func (svc *service) GetHotelName(
	condition hoteldm.GetAllHotelNameRequest,
) (results []*hotelrepo.HotelName, totalRows int64, err error) {
	return svc.repoHotelName.GetAll(condition)
}

func (svc *service) GetHotelImages(
	condition hoteldm.GetAllHotelImagesRequest,
) (results []*hotelrepo.HotelImages, totalRows int64, err error) {
	return svc.repoHotelImages.GetAll(condition)
}

func (svc *service) GetHotelFacility(
	condition hoteldm.GetAllHotelFacilityRequest,
) (results []*hotelrepo.HotelFacility, totalRows int64, err error) {
	return svc.repoHotelFacility.GetAll(condition)
}

func (svc *service) GetHotelAddress(
	condition hoteldm.GetAllHotelAddressRequest,
) (results []*hotelrepo.HotelAddress, totalRows int64, err error) {
	return svc.repoHotelAddress.GetAll(condition)
}

func (svc *service) GetHotelDescription(
	condition hoteldm.GetAllHotelDescriptionRequest,
) (results []*hotelrepo.HotelDescription, totalRows int64, err error) {
	return svc.repoHotelDescription.GetAll(condition)
}

func (svc *service) GetHotelInterestPoints(
	condition hoteldm.GetAllHotelInterestPointsRequest,
) (results []*hotelrepo.HotelInterestPoints, totalRows int64, err error) {
	return svc.repoHotelInterestPoints.GetAll(condition)
}

func (svc *service) GetHotelIssues(
	condition hoteldm.GetAllHotelIssuesRequest,
) (results []*hotelrepo.HotelIssues, totalRows int64, err error) {
	return svc.repoHotelIssues.GetAll(condition)
}

func (svc *service) GetHotelPhones(
	condition hoteldm.GetAllHotelPhoneRequest,
) (results []*hotelrepo.HotelPhone, totalRows int64, err error) {
	return svc.repoHotelPhones.GetAll(condition)
}

func (svc *service) GetRooms(
	condition hoteldm.GetAllHotelRoomsRequest,
) (results []*hotelrepo.HotelRooms, totalRows int64, err error) {
	return svc.repoHotelRooms.GetAll(condition)
}

func (svc *service) GetRoomsFacility(
	condition hoteldm.GetAllHotelRoomFacilitiesRequest,
) (results []*hotelrepo.HotelRoomFacilities, totalRows int64, err error) {
	return svc.repoHotelRoomsFacilities.GetAll(condition)
}

func (svc *service) GetRoomsStay(
	condition hoteldm.GetAllHotelRoomStayRequest,
) (results []*hotelrepo.HotelRoomStay, totalRows int64, err error) {
	return svc.repoHotelRoomsStay.GetAll(condition)
}

func (svc *service) GetCity(
	condition hoteldm.GetAllHotelCityRequest,
) (results []*hotelrepo.HotelCity, totalRows int64, err error) {
	return svc.repoHotelCity.GetAll(condition)
}

func (svc *service) GetCoordinates(
	condition hoteldm.GetAllHotelCoordinatesRequest,
) (results []*hotelrepo.HotelCoordinates, totalRows int64, err error) {
	return svc.repoCoordinates.GetAll(condition)
}
