package googlehdl

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain/googledm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servPlance ports.GooglePlaceService
}

func NewHandler(
	servPlance ports.GooglePlaceService,
) *Handler {
	return &Handler{
		servPlance: servPlance,
	}
}

func (h *Handler) GetPlaces(ctx context.Context, input model.GetPlacesInput) (*model.PlacesData, error) {
	condition := &googledm.QueryautocompleteRequest{
		Input: *input.Query,
	}

	res, err := h.servPlance.Queryautocomplete(condition)
	if err != nil {
		return nil, fmt.Errorf("ไม่พบข้อมูล")
	}

	// if res.Status != "OK" {
	// 	return nil, nil
	// }

	place := []*model.Places{}
	var predictions []googledm.Predictions
	res.BodyParser(&predictions)

	for _, v := range res.Predictions {
		n := model.Places{
			Description: v.Description,
			PlaceID:     v.PlaceID,
		}
		place = append(place, &n)
	}

	display := &model.PlacesData{
		Places: place,
	}

	return display, nil
}
