package memberhdl

import "github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"

type Handler struct {
	serv       ports.MemberService
	serviceURL string
}

func NewHandler(serv ports.MemberService, serviceURL string) *Handler {
	return &Handler{
		serv:       serv,
		serviceURL: serviceURL,
	}
}
