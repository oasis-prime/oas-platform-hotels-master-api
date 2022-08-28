package graph

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/memberhdl"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MemberHandler    *memberhdl.Handler
	HotelbedsHandler *hotelbedshdl.Handler
	HotelsHandler    *hotelshdl.Handler
}
