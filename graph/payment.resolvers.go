package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Payment is the resolver for the payment field.
func (r *mutationResolver) Payment(ctx context.Context, input model.PaymentInput) (*model.PaymentData, error) {
	return r.PaymentHandler.Payment(ctx, input)
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, input model.GetPaymentInput) (*model.PaymentData, error) {
	return r.PaymentHandler.GetPayment(ctx, input)
}
