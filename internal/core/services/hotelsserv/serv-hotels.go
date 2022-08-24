package hotelsserv

import "github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"

type Service struct {
	repoHotels ports.HotelsRepository
}

func NewService(
	repoHotels ports.HotelsRepository,
) *Service {
	return &Service{
		repoHotels: repoHotels,
	}
}
