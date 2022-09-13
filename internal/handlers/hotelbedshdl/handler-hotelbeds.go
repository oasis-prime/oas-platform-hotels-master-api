package hotelbedshdl

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"github.com/sirupsen/logrus"
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
	// condition := input.ParseToAvailabilityRequest()

	if err != nil {
		return nil, err
	}

	availability := []*model.AvailabilityHotels{}

	condition := &hotelbedsdm.AvailabilityRequest{}

	copier.CopyWithOption(&condition, &input, copier.Option{IgnoreEmpty: true})

	f, _ := json.Marshal(condition)
	logrus.Info(string(f))

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
