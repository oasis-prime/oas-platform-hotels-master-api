package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// MemberVerify is the resolver for the memberVerify field.
func (r *mutationResolver) MemberVerify(ctx context.Context, input model.MemberVerifyEmailInput) (*model.MemberVerifyEmailData, error) {
	return r.MemberHandler.GetAllHotel(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }