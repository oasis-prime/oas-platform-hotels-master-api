package handlers

import "github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"

type MemberHandler struct {
	serv       ports.MemberService
	serviceURL string
}

func NewMemberHandler(serv ports.MemberService, serviceURL string) *MemberHandler {
	return &MemberHandler{
		serv:       serv,
		serviceURL: serviceURL,
	}
}
