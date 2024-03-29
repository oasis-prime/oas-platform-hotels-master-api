package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Ticker is the resolver for the Ticker field.
func (r *mutationResolver) Ticker(ctx context.Context, input model.TickerInput) (*model.Ticker, error) {
	return r.TickerHandler.Ticker(ctx, input)
}

// GetTickers is the resolver for the getTickers field.
func (r *queryResolver) GetTickers(ctx context.Context, input model.GetTickersInput) (*model.TickerData, error) {
	return r.TickerHandler.GetTickers(ctx, input)
}

// GetTicker is the resolver for the getTicker field.
func (r *queryResolver) GetTicker(ctx context.Context, input model.GetTickerInput) (*model.Ticker, error) {
	return r.TickerHandler.GetTicker(ctx, input)
}
