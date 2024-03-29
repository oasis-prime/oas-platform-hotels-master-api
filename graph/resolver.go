package graph

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/googlehdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/memberhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/paymenthdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/popularhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/subscriptionhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/tickerhdl"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MemberHandler       *memberhdl.Handler
	HotelbedsHandler    *hotelbedshdl.Handler
	HotelsHandler       *hotelshdl.Handler
	GoogleHandler       *googlehdl.Handler
	PaymentHandler      *paymenthdl.Handler
	PopularHandler      *popularhdl.Handler
	SubscriptionHandler *subscriptionhdl.Handler
	TickerHandler       *tickerhdl.Handler
}
