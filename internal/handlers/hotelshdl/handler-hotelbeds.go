package hotelshdl

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/htenums"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	servHotels ports.HotelsService
}

func NewHandler(
	servHotels ports.HotelsService,
) *Handler {
	return &Handler{
		servHotels: servHotels,
	}
}

func (h *Handler) GetAllHotel(ctx context.Context, input model.HotelsInput) (display *model.HotelsData, err error) {
	// condition := input.ParseToAvailabilityRequest()

	hotelsRepo, err := h.servHotels.GetByCoordinates(hoteldm.GetByCoordinatesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     input.Pagination.Page,
			PageSize: input.Pagination.PageSize,
		},
		Latitude:  input.Geolocation.Latitude,
		Longitude: input.Geolocation.Longitude,
		Distance:  uint(input.Geolocation.Radius),
	})

	if err != nil {
		return nil, fmt.Errorf("")
	}

	var hotels []*model.Hotels

	for _, hotel := range hotelsRepo {
		var h model.Hotels
		copier.Copy(&h, &hotel)
		h.Language = input.Language
		hotels = append(hotels, &h)
	}

	display = &model.HotelsData{
		Hotels:     hotels,
		Pagination: &model.PaginationType{
			// Page: ,
		},
	}

	return display, nil
}

func (h *Handler) GetAllHotelName(ctx context.Context, obj *model.Hotels) (display *model.Name, err error) {
	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)
	languageType := langenums.Language(obj.Language)
	// var language langenums.Language

	hotelNames, _, err := h.servHotels.GetHotelName(hoteldm.GetAllHotelNameRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     1,
			PageSize: 1,
		},
		HotelCode: &code,
		HotelType: &hotelType,
		Language:  &languageType,
	})

	if err != nil {
		logrus.Error(err)
	}

	if len(hotelNames) > 0 {
		for _, n := range hotelNames {
			if n.Language == languageType {
				display = &model.Name{
					Content: &n.Value,
				}
				break
			}
		}

		if display == nil {
			display = &model.Name{
				Content: &hotelNames[0].Value,
			}
		}
	}

	return display, err
}

func (h *Handler) GetAllHotelImages(ctx context.Context, obj *model.Hotels) ([]*model.Images, error) {
	display := []*model.Images{}
	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)

	result, _, err := h.servHotels.GetHotelImages(hoteldm.GetAllHotelImagesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     1,
			PageSize: 20,
			Order:    "order",
		},
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		return nil, fmt.Errorf("")
	}

	for _, v := range result {
		display = append(display, &model.Images{
			ImageTypeCode: &v.ImageTypeCode,
			Path:          &v.Path,
			Order:         &v.Order,
			VisualOrder:   &v.VisualOrder,
		})
	}

	return display, nil
}

func (h *Handler) GetAllHotelFacilities(ctx context.Context, obj *model.Hotels) ([]*model.Facilities, error) {
	display := []*model.Facilities{}

	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)

	result, _, err := h.servHotels.GetHotelFacility(hoteldm.GetAllHotelFacilityRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     1,
			PageSize: 20,
			Order:    "order",
		},
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		return nil, fmt.Errorf("")
	}

	for _, v := range result {
		display = append(display, &model.Facilities{
			FacilityCode:      &v.FacilityCode,
			FacilityGroupCode: &v.FacilityGroupCode,
			Order:             &v.Order,
			Number:            &v.Number,
			Voucher:           &v.Voucher,
		})
	}

	return display, nil
}

func (h *Handler) GetAllHotelRooms(ctx context.Context, obj *model.Hotels) ([]*model.Rooms, error) {
	display := []*model.Rooms{}

	return display, nil
}

func (h *Handler) GetAllHotelPhones(ctx context.Context, obj *model.Hotels) ([]*model.Phones, error) {
	display := []*model.Phones{}

	return display, nil
}

func (h *Handler) GetAllHotelCity(ctx context.Context, obj *model.Hotels) (*model.City, error) {
	display := &model.City{}

	return display, nil
}

func (h *Handler) GetAllHotelAddress(ctx context.Context, obj *model.Hotels) (*model.Address, error) {
	display := &model.Address{}

	return display, nil
}

func (h *Handler) GetAllHotelDescription(ctx context.Context, obj *model.Hotels) (*model.Description, error) {
	display := &model.Description{}

	return display, nil
}
