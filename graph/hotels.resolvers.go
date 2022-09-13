package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
)

// Images is the resolver for the images field.
func (r *hotelResolver) Images(ctx context.Context, obj *model.Hotel, input *model.ImagesInput) ([]*model.Images, error) {
	return r.HotelsHandler.GetAllHotelImages(ctx, obj, input)
}

// InterestPoints is the resolver for the interestPoints field.
func (r *hotelResolver) InterestPoints(ctx context.Context, obj *model.Hotel, input *model.HotelInterestPointsInput) ([]*model.InterestPoints, error) {
	return r.HotelsHandler.GetAllInterestPoints(ctx, obj, input)
}

// Issues is the resolver for the issues field.
func (r *hotelResolver) Issues(ctx context.Context, obj *model.Hotel, input *model.HotelIssuesInput) ([]*model.Issues, error) {
	return r.HotelsHandler.GetAllIssues(ctx, obj, input)
}

// Facilities is the resolver for the facilities field.
func (r *hotelResolver) Facilities(ctx context.Context, obj *model.Hotel, input *model.FacilitiesInput) ([]*model.Facilities, error) {
	return r.HotelsHandler.GetAllHotelFacilities(ctx, obj, input)
}

// Rooms is the resolver for the rooms field.
func (r *hotelResolver) Rooms(ctx context.Context, obj *model.Hotel, input *model.HotelRoomsInput) ([]*model.Rooms, error) {
	return r.HotelsHandler.GetAllHotelRooms(ctx, obj, input)
}

// Phones is the resolver for the phones field.
func (r *hotelResolver) Phones(ctx context.Context, obj *model.Hotel, input *model.HotelPhonesInput) ([]*model.Phones, error) {
	return r.HotelsHandler.GetAllHotelPhones(ctx, obj, input)
}

// City is the resolver for the city field.
func (r *hotelResolver) City(ctx context.Context, obj *model.Hotel) (*model.City, error) {
	return r.HotelsHandler.GetAllHotelCity(ctx, obj)
}

// Address is the resolver for the address field.
func (r *hotelResolver) Address(ctx context.Context, obj *model.Hotel) (*model.Address, error) {
	return r.HotelsHandler.GetAllHotelAddress(ctx, obj)
}

// Description is the resolver for the description field.
func (r *hotelResolver) Description(ctx context.Context, obj *model.Hotel) (*model.Description, error) {
	return r.HotelsHandler.GetAllHotelDescription(ctx, obj)
}

// Name is the resolver for the name field.
func (r *hotelResolver) Name(ctx context.Context, obj *model.Hotel) (*model.Name, error) {
	return r.HotelsHandler.GetAllHotelName(ctx, obj)
}

// GetHotels is the resolver for the getHotels field.
func (r *queryResolver) GetHotels(ctx context.Context, input model.HotelsInput) (*model.HotelsData, error) {
	return r.HotelsHandler.GetAllHotels(ctx, input)
}

// GetHotel is the resolver for the getHotel field.
func (r *queryResolver) GetHotel(ctx context.Context, input model.HotelInput) (*model.Hotel, error) {
	return r.HotelsHandler.GetAllHotel(ctx, input)
}

// RoomStays is the resolver for the roomStays field.
func (r *roomsResolver) RoomStays(ctx context.Context, obj *model.Rooms, input *model.StaysInput) ([]*model.RoomStays, error) {
	return r.HotelsHandler.GetAllRoomStays(ctx, obj, input)
}

// RoomFacilities is the resolver for the roomFacilities field.
func (r *roomsResolver) RoomFacilities(ctx context.Context, obj *model.Rooms, input *model.FacilitiesInput) ([]*model.RoomFacilities, error) {
	return r.HotelsHandler.GetAllRoomFacilities(ctx, obj, input)
}

// RoomImages is the resolver for the roomImages field.
func (r *roomsResolver) RoomImages(ctx context.Context, obj *model.Rooms, input *model.ImagesInput) ([]*model.Images, error) {
	return r.HotelsHandler.GetAllRoomImages(ctx, obj, input)
}

// Hotel returns generated.HotelResolver implementation.
func (r *Resolver) Hotel() generated.HotelResolver { return &hotelResolver{r} }

// Rooms returns generated.RoomsResolver implementation.
func (r *Resolver) Rooms() generated.RoomsResolver { return &roomsResolver{r} }

type hotelResolver struct{ *Resolver }
type roomsResolver struct{ *Resolver }
