package hotelbedshdl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servHotelbeds ports.HotelbedsService
	servHotels    ports.HotelsService
}

func NewHandler(
	servHotelbeds ports.HotelbedsService,
) *Handler {
	return &Handler{
		servHotelbeds: servHotelbeds,
	}
}

func (h *Handler) AvailabilityHotel(ctx context.Context, input model.AvailabilityInput) (display *model.AvailabilityData, err error) {
	condition := input.ParseToAvailabilityRequest()

	if err != nil {
		return nil, err
	}

	availability := []*model.AvailabilityHotels{}

	response, err := h.servHotelbeds.AvailabilityHotelbeds(&hotelbedsdm.AvailabilityRequest{
		Stay:        condition.Stay,
		Occupancies: condition.Occupancies,
		Filter:      condition.Filter,
		Language:    condition.Language,
		Hotels:      condition.Hotels,
	})

	if err == nil {
		for _, v := range response.Hotels.Hotels {
			var a model.AvailabilityHotels
			copier.Copy(&a, &v)
			availability = append(availability, &a)
		}
	}

	display = &model.AvailabilityData{
		Availability: availability,
	}

	return display, nil
}
