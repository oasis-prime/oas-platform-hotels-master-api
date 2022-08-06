package hotelbedshdl

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servContent ports.HotelbedsContentService
}

func NewHandler(
	servContent ports.HotelbedsContentService) *Handler {
	return &Handler{
		servContent: servContent,
	}
}

func (h *Handler) CreateHotel(ctx context.Context, input model.NewHotel) (*model.Hotel, error) {

	e := h.servContent.ServiceGetAllHotelbeds()

	if e != nil {
		return nil, e
	}

	return nil, nil
}

func (h *Handler) Hotels(ctx context.Context) ([]*model.Hotel, error) {
	return nil, nil
}
