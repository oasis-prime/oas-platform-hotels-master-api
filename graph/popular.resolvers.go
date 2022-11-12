package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// GetPopular is the resolver for the getPopular field.
func (r *queryResolver) GetPopular(ctx context.Context, input model.GetPopularInput) (*model.Popular, error) {
	panic(fmt.Errorf("not implemented: GetPopular - getPopular"))
}

// GetAllPopular is the resolver for the getAllPopular field.
func (r *queryResolver) GetAllPopular(ctx context.Context, input model.GetPopularInput) (*model.PopularData, error) {
	return r.PopularHandler.GetPopular(ctx, input)
}
