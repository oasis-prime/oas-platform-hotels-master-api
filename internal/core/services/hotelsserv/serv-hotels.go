package hotelsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Service struct {
	repoHotels ports.HotelsRepository
}

func NewService(
	repoHotels ports.HotelsRepository,
) *Service {
	return &Service{
		repoHotels: repoHotels,
	}
}

func (svc *Service) GetByCoordinates(condition hoteldm.GetByCoordinatesRequest) (record []*hotelrepo.Hotels, err error) {
	return svc.repoHotels.GetByCoordinates(condition)
}
