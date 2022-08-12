package ports

import "github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"

type MemberService interface {
}

type HotelbedsService interface {
	AvailabilityHotelbeds(condition *hotelbedsdm.AvailabilityRequest) (res *hotelbedsdm.AvailabilityResponse, err error)
}
