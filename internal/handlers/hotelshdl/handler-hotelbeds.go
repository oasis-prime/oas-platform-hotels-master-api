package hotelshdl

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/htenums"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
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

func (h *Handler) GetAllHotel(
	ctx context.Context,
	input model.HotelsInput,
) (display *model.HotelsData, err error) {
	var hotelsRepo []*hotelrepo.Hotels

	hotelsRepo, total, err := h.servHotels.GetHotel(input)

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
		Hotels: hotels,
		Pagination: &model.PaginationType{
			Page:     input.Pagination.Page,
			PageSize: input.Pagination.PageSize,
			Total:    int(total),
		},
	}

	return display, nil
}

func (h *Handler) GetAllHotelName(
	ctx context.Context,
	obj *model.Hotels,
) (display *model.Name, err error) {
	if obj == nil {
		return nil, nil
	}

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
			Language: &languageType,
		},
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
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

	return display, nil
}

func (h *Handler) GetAllHotelImages(
	ctx context.Context,
	obj *model.Hotels,
	input *model.ImagesInput,
) ([]*model.Images, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.Images{}
	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)

	result, _, err := h.servHotels.GetHotelImages(hoteldm.GetAllHotelImagesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Offset:   input.Offset,
			Limit:    input.Limit,
			Order:    "order",
			IsOffset: true,
		},
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
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

func (h *Handler) GetAllHotelFacilities(
	ctx context.Context,
	obj *model.Hotels,
	input *model.FacilitiesInput,
) ([]*model.Facilities, error) {
	if obj == nil {
		return nil, nil
	}
	display := []*model.Facilities{}

	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)

	result, _, err := h.servHotels.GetHotelFacility(hoteldm.GetAllHotelFacilityRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Offset:    input.Offset,
			Limit:     input.Limit,
			Order:     "order",
			IsContent: true,
			Language:  (*langenums.Language)(&obj.Language),
			IsOffset:  true,
		},
		HotelCode:         &code,
		HotelType:         &hotelType,
		FacilityGroupCode: &input.GroupCode,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
	}

	for _, v := range result {
		display = append(display, &model.Facilities{
			FacilityName:      &v.FacilityName,
			FacilityGroupName: &v.FacilityGroupName,
			FacilityCode:      &v.FacilityCode,
			FacilityGroupCode: &v.FacilityGroupCode,
			Order:             &v.Order,
			Number:            &v.Number,
			Voucher:           &v.Voucher,
		})
	}

	return display, nil
}

func (h *Handler) GetAllHotelRooms(
	ctx context.Context,
	obj *model.Hotels,
	input *model.HotelRoomsInput,
) ([]*model.Rooms, error) {
	if obj == nil {
		return nil, nil
	}

	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	result, _, err := h.servHotels.GetRooms(hoteldm.GetAllHotelRoomsRequest{
		HotelCode: &code,
		HotelType: (*htenums.HotelTypes)(&obj.Type),
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Language:  (*langenums.Language)(&obj.Language),
			IsContent: true,
		},
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
	}

	display := []*model.Rooms{}

	for _, v := range result {
		hotelCode := int(v.HotelCode)
		display = append(display, &model.Rooms{
			RoomCode:           &v.RoomCode,
			IsParentRoom:       &v.IsParentRoom,
			MinPax:             &v.MaxPax,
			MaxPax:             &v.MaxPax,
			MaxAdults:          &v.MaxAdults,
			MaxChildren:        &v.MaxChildren,
			MinAdults:          &v.MaxAdults,
			RoomType:           &v.RoomType,
			CharacteristicCode: &v.CharacteristicCode,
			HotelCode:          &hotelCode,
			HotelType:          model.HotelTypeEnum(v.HotelType),
		})
	}

	return display, nil
}

func (h *Handler) GetAllHotelPhones(
	ctx context.Context,
	obj *model.Hotels,
	input *model.HotelPhonesInput,
) ([]*model.Phones, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.Phones{}

	return display, nil
}

func (h *Handler) GetAllHotelCity(
	ctx context.Context,
	obj *model.Hotels,
) (*model.City, error) {
	if obj == nil {
		return nil, nil
	}

	display := &model.City{}

	return display, nil
}

func (h *Handler) GetAllHotelAddress(
	ctx context.Context,
	obj *model.Hotels,
) (*model.Address, error) {
	if obj == nil {
		return nil, nil
	}

	display := &model.Address{}

	return display, nil
}

func (h *Handler) GetAllHotelDescription(
	ctx context.Context,
	obj *model.Hotels,
) (display *model.Description, err error) {
	if obj == nil {
		return nil, nil
	}

	var code uint = 0
	if obj.Code != nil {
		code = uint(*obj.Code)
	}

	hotelType := htenums.HotelTypes(obj.Type)
	languageType := langenums.Language(obj.Language)
	// var language langenums.Language

	hotelDes, _, err := h.servHotels.GetHotelDescription(hoteldm.GetAllHotelDescriptionRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Page:     1,
			PageSize: 1,
			Language: &languageType,
		},
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
	}

	if len(hotelDes) > 0 {
		for _, n := range hotelDes {
			if n.Language == languageType {
				display = &model.Description{
					Content: &n.Value,
				}
				break
			}
		}

		if display == nil {
			display = &model.Description{
				Content: &hotelDes[0].Value,
			}
		}
	}

	return display, nil
}

func (h *Handler) GetAllInterestPoints(
	ctx context.Context,
	obj *model.Hotels,
	input *model.HotelInterestPointsInput,
) ([]*model.InterestPoints, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.InterestPoints{}

	return display, nil
}

func (h *Handler) GetAllIssues(
	ctx context.Context,
	obj *model.Hotels,
	input *model.HotelIssuesInput,
) ([]*model.Issues, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.Issues{}

	return display, nil
}

func (h *Handler) GetAllRoomStays(
	ctx context.Context,
	obj *model.Rooms,
	input *model.StaysInput,
) ([]*model.RoomStays, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.RoomStays{}

	return display, nil
}

func (h *Handler) GetAllRoomFacilities(
	ctx context.Context,
	obj *model.Rooms,
	input *model.FacilitiesInput,
) ([]*model.RoomFacilities, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.RoomFacilities{}

	var code uint = 0
	hotelType := htenums.HotelTypes(obj.HotelType)
	if obj.HotelCode != nil {
		code = uint(*obj.HotelCode)
	}

	result, _, err := h.servHotels.GetRoomsFacility(hoteldm.GetAllHotelRoomFacilitiesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Offset:    input.Offset,
			Limit:     input.Limit,
			IsContent: true,
			IsOffset:  true,
		},
		FacilityGroupCode: &input.GroupCode,
		RoomCode:          obj.RoomCode,
		HotelCode:         &code,
		HotelType:         &hotelType,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
	}

	for _, v := range result {
		display = append(display, &model.RoomFacilities{
			FacilityName:      &v.FacilityName,
			FacilityCode:      &v.FacilityCode,
			FacilityGroupName: &v.FacilityGroupName,
			FacilityGroupCode: &v.FacilityGroupCode,
			IndLogic:          &v.IndLogic,
			Number:            &v.Number,
			Voucher:           &v.Voucher,
		})
	}

	return display, nil
}

func (h *Handler) GetAllRoomImages(
	ctx context.Context,
	obj *model.Rooms,
	input *model.ImagesInput,
) ([]*model.Images, error) {
	if obj == nil {
		return nil, nil
	}

	display := []*model.Images{}

	var code uint = 0
	hotelType := htenums.HotelTypes(obj.HotelType)
	if obj.HotelCode != nil {
		code = uint(*obj.HotelCode)
	}

	result, _, err := h.servHotels.GetHotelImages(hoteldm.GetAllHotelImagesRequest{
		GetAllRequestBasic: hoteldm.GetAllRequestBasic{
			Offset:   input.Offset,
			Limit:    input.Limit,
			Order:    "order",
			IsOffset: true,
		},
		RoomCode:  obj.RoomCode,
		HotelCode: &code,
		HotelType: &hotelType,
	})

	if err != nil {
		logrus.Error(err)
		return nil, nil
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
