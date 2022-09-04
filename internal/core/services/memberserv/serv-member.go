package memberserv

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	repo ports.MemberRepository
}

func NewService(repo ports.MemberRepository) *service {
	return &service{
		repo: repo,
	}
}
