package memberhdl

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-firebase-core/domain/firebasedm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"github.com/oasis-prime/oas-platform-hotels-master-api/utils"
)

type Handler struct {
	serv ports.MemberService
}

func NewHandler(serv ports.MemberService) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) VerifyEmail(
	ctx context.Context,
	input model.MemberVerifyEmailInput,
) (display *model.MemberVerifyEmailData, err error) {
	if input.Email == nil {
		return nil, fmt.Errorf("required field email")
	}

	err = h.serv.PublisherVerifyEmail(domain.PublisherVerifyEmail{
		Email: *input.Email,
		Token: "",
	})
	if err != nil {
		return nil, err
	}

	return &model.MemberVerifyEmailData{
		Message: "ok",
	}, nil
}

func (h *Handler) MemberRegister(
	ctx context.Context,
	input model.MemberRegisterInput,
) (*model.MemberRegisterData, error) {
	display := &model.MemberRegisterData{
		Message: "ok",
	}

	fmt.Printf("%v \n", input)

	err := h.serv.MemberRegister(firebasedm.MemberRegister{
		Email:       input.Email,
		PhoneNumber: "",
		Password:    input.Password,
		DisplayName: input.Display,
	})

	if err != nil {
		return nil, err
	}

	h.serv.PublisherVerifyEmail(domain.PublisherVerifyEmail{
		Email: input.Email,
		Token: utils.RandString(10),
	})

	if err != nil {
		return nil, err
	}

	return display, nil
}
