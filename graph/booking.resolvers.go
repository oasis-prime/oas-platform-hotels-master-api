package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// BookingB is the resolver for the bookingB field.
func (r *mutationResolver) BookingB(ctx context.Context, input model.BookingInput) (*model.BookingData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Booking is the resolver for the booking field.
func (r *queryResolver) Booking(ctx context.Context, input model.BookingInput) (*model.BookingData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
