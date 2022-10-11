package masterserv

import (
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"gorm.io/gorm"
)

type service struct {
	Accommodations     ports.MasterAccommodationsRepository
	Amenities          ports.MasterAmenitiesRepository
	Boards             ports.MasterBoardsRepository
	Categories         ports.MasterCategoriesRepository
	Chains             ports.MasterChainsRepository
	Currencies         ports.MasterCurrenciesRepository
	Facilities         ports.MasterFacilitiesRepository
	FacilityGroups     ports.MasterFacilityGroupsRepository
	FacilityTypologies ports.MasterFacilityTypologiesRepository
	GroupCategories    ports.MasterGroupCategoriesRepository
	HotelType          ports.MasterHotelTypeRepository
	ImageTypes         ports.MasterImageTypesRepository
	Issues             ports.MasterIssuesRepository
	Languages          ports.MasterLanguagesRepository
	Promotions         ports.MasterPromotionsRepository
	RoomBedsType       ports.MasterRoomBedsTypeRepository
	Rooms              ports.MasterRoomsRepository
	Segments           ports.MasterSegmentsRepository
	Surroundings       ports.MasterSurroundingsRepository
	Terminals          ports.MasterTerminalsRepository
	Popular            ports.MasterPopularRepository
}

func NewService(db *gorm.DB) *service {
	accommodations := masterrepo.NewAccommodationsRepo(db)
	amenities := masterrepo.NewAmenitiesRepo(db)
	boards := masterrepo.NewBoardsRepo(db)
	categories := masterrepo.NewCategoriesRepo(db)
	chains := masterrepo.NewChainsRepo(db)
	currencies := masterrepo.NewCurrenciesRepo(db)
	facilities := masterrepo.NewFacilitiesRepo(db)
	facilityGroups := masterrepo.NewFacilityGroupsRepo(db)
	facilityTypologies := masterrepo.NewFacilityTypologiesRepo(db)
	groupCategories := masterrepo.NewGroupCategoriesRepo(db)
	hotelType := masterrepo.NewHotelTypeRepo(db)
	imageTypes := masterrepo.NewImageTypesRepo(db)
	issues := masterrepo.NewIssuesRepo(db)
	languages := masterrepo.NewLanguagesRepo(db)
	promotions := masterrepo.NewPromotionsRepo(db)
	roomBeds := masterrepo.NewRoomBedsTypeRepo(db)
	rooms := masterrepo.NewRoomsRepo(db)
	segments := masterrepo.NewSegmentsRepo(db)
	surroundings := masterrepo.NewSurroundingsRepo(db)
	terminals := masterrepo.NewTerminalsRepo(db)
	popular := masterrepo.NewPopularRepo(db)

	return &service{
		Accommodations:     accommodations,
		Amenities:          amenities,
		Boards:             boards,
		Categories:         categories,
		Chains:             chains,
		Currencies:         currencies,
		Facilities:         facilities,
		FacilityGroups:     facilityGroups,
		FacilityTypologies: facilityTypologies,
		GroupCategories:    groupCategories,
		HotelType:          hotelType,
		ImageTypes:         imageTypes,
		Issues:             issues,
		Languages:          languages,
		Promotions:         promotions,
		RoomBedsType:       roomBeds,
		Rooms:              rooms,
		Segments:           segments,
		Surroundings:       surroundings,
		Terminals:          terminals,
		Popular:            popular,
	}
}
