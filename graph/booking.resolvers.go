package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Booking is the resolver for the booking field.
func (r *mutationResolver) Booking(ctx context.Context, input model.BookingInput) (*model.BookingData, error) {
	return r.PaymentHandler.Booking(ctx, input)
}

// GetBooking is the resolver for the getBooking field.
func (r *queryResolver) GetBooking(ctx context.Context, input model.GetBookingInput) (*model.BookingData, error) {
	panic(fmt.Errorf("not implemented: GetBooking - getBooking"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
