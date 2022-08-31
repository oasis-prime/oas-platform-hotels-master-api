package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Images is the resolver for the images field.
func (r *hotelsResolver) Images(ctx context.Context, obj *model.Hotels) ([]*model.Images, error) {
	return r.HotelsHandler.GetAllHotelImages(ctx, obj)
}

// InterestPoints is the resolver for the interestPoints field.
func (r *hotelsResolver) InterestPoints(ctx context.Context, obj *model.Hotels) ([]*model.InterestPoints, error) {
	panic(fmt.Errorf("not implemented: InterestPoints - interestPoints"))
}

// Issues is the resolver for the issues field.
func (r *hotelsResolver) Issues(ctx context.Context, obj *model.Hotels) ([]*model.Issues, error) {
	panic(fmt.Errorf("not implemented: Issues - issues"))
}

// Facilities is the resolver for the facilities field.
func (r *hotelsResolver) Facilities(ctx context.Context, obj *model.Hotels) ([]*model.Facilities, error) {
	return r.HotelsHandler.GetAllHotelFacilities(ctx, obj)
}

// Rooms is the resolver for the rooms field.
func (r *hotelsResolver) Rooms(ctx context.Context, obj *model.Hotels) ([]*model.Rooms, error) {
	return r.HotelsHandler.GetAllHotelRooms(ctx, obj)
}

// Phones is the resolver for the phones field.
func (r *hotelsResolver) Phones(ctx context.Context, obj *model.Hotels) ([]*model.Phones, error) {
	return r.HotelsHandler.GetAllHotelPhones(ctx, obj)
}

// City is the resolver for the city field.
func (r *hotelsResolver) City(ctx context.Context, obj *model.Hotels) (*model.City, error) {
	return r.HotelsHandler.GetAllHotelCity(ctx, obj)
}

// Address is the resolver for the address field.
func (r *hotelsResolver) Address(ctx context.Context, obj *model.Hotels) (*model.Address, error) {
	return r.HotelsHandler.GetAllHotelAddress(ctx, obj)
}

// Description is the resolver for the description field.
func (r *hotelsResolver) Description(ctx context.Context, obj *model.Hotels) (*model.Description, error) {
	return r.HotelsHandler.GetAllHotelDescription(ctx, obj)
}

// Name is the resolver for the name field.
func (r *hotelsResolver) Name(ctx context.Context, obj *model.Hotels) (*model.Name, error) {
	return r.HotelsHandler.GetAllHotelName(ctx, obj)
}

// GetHotels is the resolver for the getHotels field.
func (r *queryResolver) GetHotels(ctx context.Context, input model.HotelsInput) (*model.HotelsData, error) {
	return r.HotelsHandler.GetAllHotel(ctx, input)
}

// Hotels returns generated.HotelsResolver implementation.
func (r *Resolver) Hotels() generated.HotelsResolver { return &hotelsResolver{r} }

type hotelsResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *queryResolver) Hotels(ctx context.Context, input model.HotelsInput) (*model.HotelsData, error) {
// 	return r.HotelsHandler.GetAllHotel(ctx, input)
// }
