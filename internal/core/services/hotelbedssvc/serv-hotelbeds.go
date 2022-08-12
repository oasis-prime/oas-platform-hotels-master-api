package hotelbedssvc

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Service struct {
	httpContent ports.HotelbedsHttp
	repoHotels  ports.HotelsRepository
}

func NewService(
	httpContent ports.HotelbedsHttp,
	repoHotels ports.HotelsRepository,
) *Service {
	return &Service{
		httpContent: httpContent,
		repoHotels:  repoHotels,
	}
}

func (svc *Service) AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error) {
	res, err = svc.httpContent.GetAvailability(condition)
	return res, err
}
