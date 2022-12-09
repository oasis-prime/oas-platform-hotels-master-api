package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// GetSubscription is the resolver for the getSubscription field.
func (r *queryResolver) GetSubscription(ctx context.Context, input model.SubscriptionInput) (*model.SubscriptionData, error) {
	return r.SubscriptionHandler.GetSubscription(ctx, input)
}