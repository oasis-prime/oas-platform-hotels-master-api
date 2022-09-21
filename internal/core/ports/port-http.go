package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
)

type HotelbedsHttp interface {
	// Content
	GetHotels(condition *hotelbedsdm.HotelsRequest) (response *hotelbedsdm.HotelsResponse, err error)

	// Booking
	GetAvailability(condition *hotelbedsdm.AvailabilityRequest) (response *hotelbedsdm.AvailabilityResponse, err error)
	GetCheckRate(condition *hotelbedsdm.CheckRatesRequest) (response *hotelbedsdm.CheckRatesResponse, err error)
}

type PaymentChillpayHttp interface {
	GetPaylinkGenerate(condition *chillpaydm.PaylinkGenerateRequest) (response *chillpaydm.PaylinkGenerateResponse, err error)
	GetPaylinkDetail(condition *chillpaydm.PaylinkDetailsRequest) (response *chillpaydm.PaylinkGenerateResponse, err error)
}
