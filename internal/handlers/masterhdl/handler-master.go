package masterhdl

import "github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"

type Handler struct {
	servMaster ports.MasterService
}

func NewHandler(
	servMaster ports.MasterService,
) *Handler {
	return &Handler{
		servMaster: servMaster,
	}
}
