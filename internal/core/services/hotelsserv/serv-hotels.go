package hotelsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Service struct {
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
) *Service {
	return &Service{
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

func (svc *Service) GetByCoordinates(condition hoteldm.GetByCoordinatesRequest) (record []*hotelrepo.Hotels, err error) {
	return svc.repoHotels.GetByCoordinates(condition)
}

func (svc *Service) GetHotelName(condition hoteldm.GetAllHotelNameRequest) (results []*hotelrepo.HotelName, totalRows int64, err error) {
	return svc.repoHotelName.GetAll(condition)
}

func (svc *Service) GetHotelImages(condition hoteldm.GetAllHotelImagesRequest) (results []*hotelrepo.HotelImages, totalRows int64, err error) {
	return svc.repoHotelImages.GetAll(condition)
}

func (svc *Service) GetHotelFacility(condition hoteldm.GetAllHotelFacilityRequest) (results []*hotelrepo.HotelFacility, totalRows int64, err error) {
	return svc.repoHotelFacility.GetAll(condition)
}
