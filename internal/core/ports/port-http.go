package ports

import "github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"

type HotelbedsHttp interface {
	// Content
	GetHotels(condition *hotelbedsdm.HotelsRequest) (response *hotelbedsdm.HotelsResponse, err error)

	// Booking
	GetAvailability(condition *hotelbedsdm.AvailabilityRequest) (response *hotelbedsdm.AvailabilityResponse, err error)
}
