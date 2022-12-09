package subscriptionhdl

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servSubscription ports.SubscriptionService
}

func NewHandler(
	servSubscription ports.SubscriptionService,
) *Handler {
	return &Handler{
		servSubscription: servSubscription,
	}
}

func (h *Handler) GetSubscription(ctx context.Context, input model.SubscriptionInput) (display *model.SubscriptionData, err error) {
	return display, err
}
