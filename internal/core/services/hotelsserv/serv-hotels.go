package hotelsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	repoHotels              ports.HotelsRepository
	repoHotelName           ports.HotelNameRepository
	repoHotelDescription    ports.HotelDescriptionRepository
	repoHotelInterestPoints ports.HotelInterestPointsRepository
	repoHotelIssues         ports.HotelIssuesRepository
	repoHotelFacility       ports.HotelFacilityRepository
	repoHotelRooms          ports.HotelRoomsRepository
	repoHotelPhones         ports.HotelPhonesRepository
	repoHotelCity           ports.HotelCityRepository
	repoHotelAddress        ports.HotelAddressRepository
	repoHotelImages         ports.HotelImagesRepository
}

func NewService(
	repoHotels ports.HotelsRepository,
	repoHotelName ports.HotelNameRepository,
	repoHotelDescription ports.HotelDescriptionRepository,
	repoHotelInterestPoints ports.HotelInterestPointsRepository,
	repoHotelIssues ports.HotelIssuesRepository,
	repoHotelFacility ports.HotelFacilityRepository,
	repoHotelRooms ports.HotelRoomsRepository,
	repoHotelPhones ports.HotelPhonesRepository,
	repoHotelCity ports.HotelCityRepository,
	repoHotelAddress ports.HotelAddressRepository,
	repoHotelImages ports.HotelImagesRepository,
) *service {
	return &service{
		repoHotels:              repoHotels,
		repoHotelName:           repoHotelName,
		repoHotelDescription:    repoHotelDescription,
		repoHotelInterestPoints: repoHotelInterestPoints,
		repoHotelIssues:         repoHotelIssues,
		repoHotelFacility:       repoHotelFacility,
		repoHotelRooms:          repoHotelRooms,
		repoHotelPhones:         repoHotelPhones,
		repoHotelCity:           repoHotelCity,
		repoHotelAddress:        repoHotelAddress,
		repoHotelImages:         repoHotelImages,
	}
}

func (svc *service) GetHotel(input model.HotelsInput) (record []*hotelrepo.Hotels, totalRows int64, err error) {
	if input.Geolocation != nil {
		return svc.repoHotels.GetByCoordinates(hoteldm.GetByCoordinatesRequest{
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				Page:     input.Pagination.Page,
				PageSize: input.Pagination.PageSize,
			},
			Latitude:  input.Geolocation.Latitude,
			Longitude: input.Geolocation.Longitude,
			Distance:  uint(input.Geolocation.Radius),
		})
	}

	if input.Keywords != nil {
		return svc.repoHotels.GetAll(hoteldm.GetAllRequestBasic{
			Page:     input.Pagination.Page,
			PageSize: input.Pagination.PageSize,
			Keyword:  &input.Keywords.Keyword,
		})
	}

	return nil, 0, nil
}

func (svc *service) GetHotelName(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error) {
	return svc.repoHotelName.GetAll(condition)
}

func (svc *service) GetHotelImages(condition hoteldm.GetAllHotelImagesRequest) (results []*hotelrepo.HotelImages, totalRows int64, err error) {
	return svc.repoHotelImages.GetAll(condition)
}

func (svc *service) GetHotelFacility(condition hoteldm.GetAllHotelFacilityRequest) (results []*hotelrepo.HotelFacility, totalRows int64, err error) {
	return svc.repoHotelFacility.GetAll(condition)
}
