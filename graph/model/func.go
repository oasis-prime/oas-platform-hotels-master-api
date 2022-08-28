package model

import (
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
)

func (input AvailabilityInput) ParseToAvailabilityRequest() *hotelbedsdm.AvailabilityRequest {
	occ := []hotelbedsdm.OccupanciesReq{}
	{
		for _, v := range input.Occupancies {
			occ = append(occ, hotelbedsdm.OccupanciesReq{
				Rooms:    int32(v.Rooms),
				Adults:   int32(v.Adults),
				Children: int32(v.Children),
			})
		}
	}
	// radius := uint(input.Geolocation.Radius)

	var filter *hotelbedsdm.FilterReq
	{

		if input.Filter != nil {
			var maxHotels, maxRooms, maxRatesPerRoom int32
			if input.Filter.MaxHotels != nil {
				maxHotels = int32(*input.Filter.MaxHotels)
			}
			if input.Filter.MaxRooms != nil {
				maxRooms = int32(*input.Filter.MaxRooms)
			}
			var minRate, maxRate uint
			if input.Filter.MinRate != nil {
				minRate = uint(*input.Filter.MinRate)
			}

			if input.Filter.MaxRate != nil {
				maxRate = uint(*input.Filter.MaxRate)
			}

			filter = &hotelbedsdm.FilterReq{
				MaxHotels:       &maxHotels,
				MaxRooms:        &maxRooms,
				MinRate:         &minRate,
				MaxRate:         &maxRate,
				MaxRatesPerRoom: &maxRatesPerRoom,
			}
		}
	}

	return &hotelbedsdm.AvailabilityRequest{
		Stay: hotelbedsdm.StayReq{
			CheckIn:  input.Stay.CheckIn,
			CheckOut: input.Stay.CheckOut,
		},
		Occupancies: occ,
		Hotels: &hotelbedsdm.HotelsReq{
			input.Hotels.Hotel,
		},
		Language: hotelbedsdm.Language(input.Language),
		Filter:   filter,
	}
}
