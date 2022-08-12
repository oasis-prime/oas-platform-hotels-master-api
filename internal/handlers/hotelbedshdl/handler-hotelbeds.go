package hotelbedshdl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servContent ports.HotelbedsService
}

func NewHandler(
	servContent ports.HotelbedsService) *Handler {
	return &Handler{
		servContent: servContent,
	}
}

func (h *Handler) AvailabilityHotel(ctx context.Context, input model.AvailabilityInput) (display *model.AvailabilityData, err error) {
	condition := input.ParseToAvailabilityRequest()

	response, e := h.servContent.AvailabilityHotelbeds(condition)

	if e != nil {
		return nil, e
	}

	hotels := []*model.AvailabilityHotels{}

	for _, v := range response.Hotels.Hotels {
		var a model.AvailabilityHotels
		copier.Copy(&a, &v)
		hotels = append(hotels, &a)
	}
	display = &model.AvailabilityData{
		Hotels: hotels,
	}

	return display, nil
}
