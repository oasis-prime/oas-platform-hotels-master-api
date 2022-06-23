package services

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type MemberService struct {
	repo ports.MemberRepository
}

func NewMemberService(repo ports.MemberRepository) *MemberService {
	return &MemberService{
		repo: repo,
	}
}
