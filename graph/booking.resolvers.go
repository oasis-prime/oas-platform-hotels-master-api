package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Booking is the resolver for the booking field.
func (r *mutationResolver) Booking(ctx context.Context, input model.BookingInput) (*model.Booking, error) {
	return r.PaymentHandler.Booking(ctx, input)
}

// GetBooking is the resolver for the getBooking field.
func (r *queryResolver) GetBooking(ctx context.Context, input model.GetBookingInput) (*model.Booking, error) {
	panic(fmt.Errorf("not implemented: GetBooking - getBooking"))
}

// GetBookingsHistory is the resolver for the getBookingsHistory field.
func (r *queryResolver) GetBookingsHistory(ctx context.Context, input model.BookingsHistoryInput) (*model.BookingData, error) {
	panic(fmt.Errorf("not implemented: GetBookingsHistory - getBookingsHistory"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
