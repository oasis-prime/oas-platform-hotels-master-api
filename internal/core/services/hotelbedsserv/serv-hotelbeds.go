package hotelbedsserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	httpContent ports.HotelbedsHttp
}

func NewService(
	httpContent ports.HotelbedsHttp,
) *service {
	return &service{
		httpContent: httpContent,
	}
}

func (svc *service) AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error) {
	res, err = svc.httpContent.GetAvailability(condition)
	return res, err
}

func (svc *service) CheckRate(condition *hotelbedsdm.CheckRatesRequest) (res *hotelbedsdm.CheckRatesResponse, err error) {
	res, err = svc.httpContent.GetCheckRate(condition)
	return res, err
}
