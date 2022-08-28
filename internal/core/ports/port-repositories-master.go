package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
)

type MasterAccommodationsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllAccommodationsDescriptionRequest) (results []*masterrepo.AccommodationsDescription, totalRows int64, err error)
}

type MasterAccommodationsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Accommodations, totalRows int64, err error)
}

type MasterAmenitiesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Amenities, totalRows int64, err error)
}

type MasterBoardsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllBoardsDescriptionRequest) (results []*masterrepo.BoardsDescription, totalRows int64, err error)
}

type MasterBoardsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Boards, totalRows int64, err error)
}

type MasterCategoriesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllCategoriesDescriptionRequest) (results []*masterrepo.CategoriesDescription, totalRows int64, err error)
}

type MasterCategoriesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Categories, totalRows int64, err error)
}

type MasterChainsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllChainsDescriptionRequest) (results []*masterrepo.ChainsDescription, totalRows int64, err error)
}

type MasterChainsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Chains, totalRows int64, err error)
}

type MasterCurrenciesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllCurrenciesDescriptionRequest) (results []*masterrepo.CurrenciesDescription, totalRows int64, err error)
}

type MasterCurrenciesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Currencies, totalRows int64, err error)
}

type MasterFacilitiesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllFacilitiesDescriptionRequest) (results []*masterrepo.FacilitiesDescription, totalRows int64, err error)
}

type MasterFacilityTypologiesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.FacilityTypologies, totalRows int64, err error)
}

type MasterFacilityGroupsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllFacilityGroupsDescriptionRequest) (results []*masterrepo.FacilityGroupsDescription, totalRows int64, err error)
}

type MasterFacilityGroupsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.FacilityGroups, totalRows int64, err error)
}

type MasterFacilitiesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Facilities, totalRows int64, err error)
}

type MasterGroupCategoriesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllGroupCategoriesDescriptionRequest) (results []*masterrepo.GroupCategoriesDescription, totalRows int64, err error)
}

type MasterGroupCategoriesNameRepository interface {
	GetAll(condition masterdm.GetAllGroupCategoriesNameRequest) (results []*masterrepo.GroupCategoriesName, totalRows int64, err error)
}

type MasterGroupCategoriesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.GroupCategories, totalRows int64, err error)
}

type MasterHotelTypeRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.HotelType, totalRows int64, err error)
}

type MasterImageTypeDescriptionRepository interface {
	GetAll(condition masterdm.GetAllImageTypeDescriptionRequest) (results []*masterrepo.ImageTypeDescription, totalRows int64, err error)
}

type MasterImageTypesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.ImageTypes, totalRows int64, err error)
}

type MasterIssuesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllIssuesDescriptionRequest) (results []*masterrepo.IssuesDescription, totalRows int64, err error)
}

type MasterIssuesNameRepository interface {
	GetAll(condition masterdm.GetAllIssuesNameRequest) (results []*masterrepo.IssuesName, totalRows int64, err error)
}

type MasterIssuesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Issues, totalRows int64, err error)
}

type MasterLanguagesDescriptionRepository interface {
	GetAll(condition masterdm.GetAllLanguagesDescriptionRequest) (results []*masterrepo.LanguagesDescription, totalRows int64, err error)
}

type MasterLanguagesRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Languages, totalRows int64, err error)
}

type MasterPromotionsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllPromotionsDescriptionRequest) (results []*masterrepo.PromotionsDescription, totalRows int64, err error)
}

type MasterPromotionsNameRepository interface {
	GetAll(condition masterdm.GetAllPromotionsNameRequest) (results []*masterrepo.PromotionsName, totalRows int64, err error)
}

type MasterPromotionsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Promotions, totalRows int64, err error)
}

type MasterRoomBedsTypeRepository interface {
	GetAll(condition masterdm.GetAllRoomBedsTypeRequest) (results []*masterrepo.RoomBedsType, totalRows int64, err error)
}

type MasterRoomsCharacteristicDescriptionRepository interface {
	GetAll(condition masterdm.GetAllRoomsCharacteristicDescriptionRequest) (results []*masterrepo.RoomsCharacteristicDescription, totalRows int64, err error)
}

type MasterRoomsTypeDescriptionRepository interface {
	GetAll(condition masterdm.GetAllRoomsTypeDescriptionRequest) (results []*masterrepo.RoomsTypeDescription, totalRows int64, err error)
}

type MasterRoomsRepository interface {
	GetAll(condition masterdm.GetAllRoomsRequest) (results []*masterrepo.Rooms, totalRows int64, err error)
}

type MasterSegmentsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllSegmentsDescriptionRequest) (results []*masterrepo.SegmentsDescription, totalRows int64, err error)
}

type MasterSegmentsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Segments, totalRows int64, err error)
}

type MasterSurroundingsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Surroundings, totalRows int64, err error)
}

type MasterTerminalsDescriptionRepository interface {
	GetAll(condition masterdm.GetAllTerminalsDescriptionRequest) (results []*masterrepo.TerminalsDescription, totalRows int64, err error)
}

type MasterTerminalsNameRepository interface {
	GetAll(condition masterdm.GetAllTerminalsNameRequest) (results []*masterrepo.TerminalsName, totalRows int64, err error)
}

type MasterTerminalsRepository interface {
	GetAll(condition masterdm.GetAllBasicRequest) (results []*masterrepo.Terminals, totalRows int64, err error)
}
