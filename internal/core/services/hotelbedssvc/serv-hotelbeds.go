package hotelbedssvc

import (
	"encoding/json"

	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/repositories"
	"github.com/oasis-prime/oas-platform-core/repositories/enums"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"gorm.io/datatypes"
)

type Service struct {
	httpContent ports.HotelbedsContentHttp
	repo        ports.HotelbedsContentRepository
}

func NewService(
	httpContent ports.HotelbedsContentHttp,
	repo ports.HotelbedsContentRepository) *Service {
	return &Service{
		httpContent: httpContent,
		repo:        repo,
	}
}

func (svc *Service) ServiceGetAllHotelbeds() error {
	condition := &hotelbedsdm.HotelsRequest{
		Fields:               "all",
		Language:             "ENG", //ENG | TAI
		From:                 1,
		To:                   1,
		UseSecondaryLanguage: false,
	}
	response, err := svc.httpContent.GetHotels(condition)
	if err != nil {
		return err
	}

	objHotels := []hotelbedsdm.Hotels{}
	response.BodyParser(&objHotels)

	hotels := []repositories.Hotels{}
	for _, h := range objHotels {
		var segmentCodes, boardCodes, amenityCodes []byte
		{
			segmentCodes, _ = json.Marshal(h.SegmentCodes)
			boardCodes, _ = json.Marshal(h.BoardCodes)
			amenityCodes, _ = json.Marshal(h.AmenityCodes)
		}

		var rooms []repositories.HotelRooms
		{
			for _, r := range h.Rooms {
				rooms = append(rooms, repositories.HotelRooms{
					RoomCode:           &r.RoomCode,
					IsParentRoom:       r.IsParentRoom,
					MinPax:             &r.MaxPax,
					MaxPax:             &r.MaxPax,
					MaxAdults:          &r.MaxAdults,
					MaxChildren:        &r.MaxChildren,
					MinAdults:          &r.MinAdults,
					RoomType:           &r.RoomType,
					CharacteristicCode: &r.CharacteristicCode,
				})
			}
		}

		var phones []repositories.HotelPhone
		{
			for _, p := range h.Phones {
				phones = append(phones, repositories.HotelPhone{
					PhoneNumber: &p.PhoneNumber,
					PhoneType:   &p.PhoneType,
				})
			}
		}

		var facilities []repositories.HotelFacility
		{
			for _, f := range h.Facilities {
				facilities = append(facilities, repositories.HotelFacility{
					FacilityCode:      &f.FacilityCode,
					FacilityGroupCode: &f.FacilityGroupCode,
					Order:             &f.Order,
					IndYesOrNo:        &f.IndYesOrNo,
					Number:            &f.Number,
					Voucher:           &f.Voucher,
					IndLogic:          &f.IndLogic,
					IndFee:            &f.IndFee,
					Distance:          &f.Distance,
					TimeFrom:          &f.TimeFrom,
					TimeTo:            &f.TimeTo,
					DateTo:            &f.DateTo,
				})
			}
		}

		var issues []repositories.HotelIssues
		{
			for _, is := range h.Issues {
				issues = append(issues, repositories.HotelIssues{
					IssueCode:   &is.IssueCode,
					IssueType:   &is.IssueType,
					DateFrom:    &is.DateFrom,
					DateTo:      &is.DateTo,
					Order:       &is.Order,
					Alternative: &is.Alternative,
				})
			}
		}

		var images []repositories.HotelImages
		{
			for _, ima := range h.Images {
				images = append(images, repositories.HotelImages{
					ImageTypeCode:      &ima.ImageTypeCode,
					Path:               &ima.Path,
					Order:              &ima.Order,
					VisualOrder:        &ima.VisualOrder,
					RoomCode:           &ima.RoomCode,
					RoomType:           &ima.RoomType,
					CharacteristicCode: &ima.CharacteristicCode,
				})
			}
		}

		var interestPoints []repositories.HotelInterestPoints
		{
			for _, inte := range h.InterestPoints {
				interestPoints = append(interestPoints, repositories.HotelInterestPoints{
					FacilityCode:      &inte.FacilityCode,
					FacilityGroupCode: &inte.FacilityGroupCode,
					Order:             &inte.Order,
					PoiName:           &inte.PoiName,
					Distance:          &inte.Distance,
				})
			}
		}

		hotels = append(hotels, repositories.Hotels{
			Code: uint(h.Code),
			Type: enums.Hotelbeds,
			Name: []repositories.HotelName{
				{
					Language: enums.English,
					Value:    h.Name.Content,
				},
			},
			Description: []repositories.HotelDescription{
				{
					Language: enums.English,
					Value:    h.Description.Content,
				},
			},
			CountryCode:     &h.CountryCode,
			StateCode:       &h.StateCode,
			DestinationCode: &h.DestinationCode,
			ZoneCode:        &h.ZoneCode,
			Coordinates: repositories.HotelCoordinates{
				Longitude: h.Coordinates.Longitude,
				Latitude:  h.Coordinates.Latitude,
			},
			CategoryCode:          &h.CategoryCode,
			CategoryGroupCode:     &h.CategoryGroupCode,
			AccommodationTypeCode: &h.AccommodationTypeCode,
			BoardCodes:            datatypes.JSON(boardCodes),
			SegmentCodes:          datatypes.JSON(segmentCodes),
			ChainCode:             &h.ChainCode,
			AmenityCodes:          amenityCodes,
			Address: repositories.HotelAddress{
				Language: enums.English,
				Content:  h.Address.Content,
				Street:   h.Address.Street,
				Number:   h.Address.Number,
			},
			PostalCode: &h.PostalCode,
			City: repositories.HotelCity{
				Language: enums.English,
			},
			Email:          &h.Email,
			Phones:         phones,
			Rooms:          rooms,
			Facilities:     facilities,
			Issues:         issues,
			Images:         images,
			InterestPoints: interestPoints,
			Web:            &h.Web,
			Ranking:        uint(h.Ranking),
		})
	}

	_, err = svc.repo.CreateBatch(&hotels)

	if err != nil {
		return err
	}

	return nil
}
