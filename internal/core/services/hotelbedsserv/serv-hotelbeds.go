package hotelbedsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Service struct {
	httpContent ports.HotelbedsHttp
}

func NewService(
	httpContent ports.HotelbedsHttp,
) *Service {
	return &Service{
		httpContent: httpContent,
	}
}

func (svc *Service) AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error) {
	res, err = svc.httpContent.GetAvailability(condition)
	return res, err
}
