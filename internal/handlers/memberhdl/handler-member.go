package memberhdl

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	serv ports.MemberService
}

func NewHandler(serv ports.MemberService) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) GetAllHotel(
	ctx context.Context,
	input model.MemberVerifyEmailInput,
) (display *model.MemberVerifyEmailData, err error) {
	err = h.serv.PublisherVerifyEmail()
	if err != nil {
		return nil, err
	}

	return &model.MemberVerifyEmailData{
		Message: "ok",
	}, nil
}
