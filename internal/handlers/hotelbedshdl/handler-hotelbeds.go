package hotelbedshdl

import (
	"context"
	"fmt"

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
	fmt.Println("A")

	h.servContent.ServiceGetAllHotelbeds()

	return nil, nil
}

func (h *Handler) Hotels(ctx context.Context) ([]*model.Hotel, error) {
	return nil, nil
}
