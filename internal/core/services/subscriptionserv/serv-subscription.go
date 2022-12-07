package subscriptionserv

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	repo ports.SubscriptionRepository
}

func NewService(repo ports.SubscriptionRepository) *service {
	return &service{
		repo: repo,
	}
}
