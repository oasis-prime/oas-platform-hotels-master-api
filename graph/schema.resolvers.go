package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// CreateHotel is the resolver for the createHotel field.
func (r *mutationResolver) CreateHotel(ctx context.Context, input model.NewHotel) (*model.Hotel, error) {
	return r.HotelbedsHandler.CreateHotel(ctx, input)
}

// Hotels is the resolver for the hotels field.
func (r *queryResolver) Hotels(ctx context.Context) ([]*model.Hotel, error) {
	return r.HotelbedsHandler.Hotels(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
