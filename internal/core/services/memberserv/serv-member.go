package memberserv

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Service struct {
	repo ports.MemberRepository
}

func NewService(repo ports.MemberRepository) *Service {
	return &Service{
		repo: repo,
	}
}
