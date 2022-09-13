package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// GetAvailability is the resolver for the getAvailability field.
func (r *queryResolver) GetAvailability(ctx context.Context, input model.AvailabilityInput) (*model.AvailabilityData, error) {
	return r.HotelbedsHandler.AvailabilityHotel(ctx, input)
}
