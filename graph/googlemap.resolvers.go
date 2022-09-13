package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// GetPlaces is the resolver for the getPlaces field.
func (r *queryResolver) GetPlaces(ctx context.Context, input model.GetPlacesInput) (*model.PlacesData, error) {
	return r.GoogleHandler.GetPlaces(ctx, input)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
