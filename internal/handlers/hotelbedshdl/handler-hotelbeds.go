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
}

func NewHandler(
	servHotelbeds ports.HotelbedsService,
) *Handler {
	return &Handler{
		servHotelbeds: servHotelbeds,
	}
}

func (h *Handler) AvailabilityHotel(ctx context.Context, input model.AvailabilityInput) (display *model.AvailabilityData, err error) {
	if err != nil {
		return nil, err
	}

	availability := []*model.AvailabilityHotels{}

	condition := &hotelbedsdm.AvailabilityRequest{}

	copier.CopyWithOption(&condition, &input, copier.Option{IgnoreEmpty: true})

	response, err := h.servHotelbeds.AvailabilityHotelbeds(condition)

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

func (h *Handler) CheckRate(ctx context.Context, input model.CheckRateInput) (*model.CheckRateData, error) {
	display := model.CheckRateData{}

	result, err := h.servHotelbeds.CheckRate(&hotelbedsdm.CheckRatesRequest{
		Rooms: []hotelbedsdm.RateKey{
			{RateKey: input.RateKey},
		},
		Language: hotelbedsdm.Language(input.Language),
	})

	if err != nil {
		return nil, err
	}

	copier.Copy(&display, &result.Hotel)

	return &display, nil
}
