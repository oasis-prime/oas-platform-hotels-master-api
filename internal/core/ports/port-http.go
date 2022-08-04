package ports

import "github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"

type HotelbedsContentHttp interface {
	GetHotels(condition *hotelbedsdm.HotelsRequest) (response *hotelbedsdm.HotelsResponse, err error)
}
