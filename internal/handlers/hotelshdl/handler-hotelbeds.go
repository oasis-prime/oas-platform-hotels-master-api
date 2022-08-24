package hotelshdl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servHotelbeds ports.HotelbedsService
	servHotels    ports.HotelsService
}

func NewHandler(
	servHotelbeds ports.HotelbedsService,
	servHotels ports.HotelsService,
) *Handler {
	return &Handler{
		servHotelbeds: servHotelbeds,
		servHotels:    servHotels,
	}
}

func (h *Handler) AvailabilityHotel(ctx context.Context, input model.AvailabilityInput) (display *model.AvailabilityData, err error) {
	condition := input.ParseToAvailabilityRequest()

	hotelsRepo, err := h.servHotels.GetByCoordinates(hoteldm.GetByCoordinatesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     1,
			PageSize: 20,
		},
		Latitude:  input.Geolocation.Latitude,
		Longitude: input.Geolocation.Longitude,
		Distance:  *condition.Geolocation.Radius,
	})

	if err != nil {
		return nil, err
	}

	var hotelIds []int
	var hotels []*model.Hotels
	availability := []*model.AvailabilityHotels{}
	for _, v := range hotelsRepo {
		if v.Code > 0 {
			hotelIds = append(hotelIds, int(v.Code))
		}

		hotel := &model.Hotels{}
		if v != nil {
			hotel.ParseModel(*v)
			hotels = append(hotels, hotel)
		}
	}

	response, err := h.servHotelbeds.AvailabilityHotelbeds(&hotelbedsdm.AvailabilityRequest{
		Stay:        condition.Stay,
		Occupancies: condition.Occupancies,
		Filter:      condition.Filter,
		Language:    condition.Language,
		Hotels: &hotelbedsdm.HotelsReq{
			Hotel: hotelIds,
		},
	})

	if err == nil {
		for _, v := range response.Hotels.Hotels {
			var a model.AvailabilityHotels
			copier.Copy(&a, &v)
			availability = append(availability, &a)
		}
	}

	display = &model.AvailabilityData{
		Hotels:       hotels,
		Availability: availability,
	}

	return display, nil
}
