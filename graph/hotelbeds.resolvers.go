package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// GetAvailability is the resolver for the getAvailability field.
func (r *queryResolver) GetAvailability(ctx context.Context, input model.AvailabilityInput) (*model.AvailabilityData, error) {
	panic(fmt.Errorf("not implemented: GetAvailability - getAvailability"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Availability(ctx context.Context, input model.AvailabilityInput) (*model.AvailabilityData, error) {
	return r.HotelbedsHandler.AvailabilityHotel(ctx, input)
}
