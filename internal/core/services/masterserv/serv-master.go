package masterserv

import (
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
	"gorm.io/gorm"
)

type service struct {
	repoMasterACD  ports.MasterAccommodationsDescriptionRepository
	repoMasterAC   ports.MasterAccommodationsRepository
	repoMasterAM   ports.MasterAmenitiesRepository
	repoMasterBOD  ports.MasterBoardsDescriptionRepository
	repoMasterBO   ports.MasterBoardsRepository
	repoMasterCAD  ports.MasterCategoriesDescriptionRepository
	repoMasterCA   ports.MasterCategoriesRepository
	repoMasterCHD  ports.MasterChainsDescriptionRepository
	repoMasterCH   ports.MasterChainsRepository
	repoMasterCUD  ports.MasterCurrenciesDescriptionRepository
	repoMasterCU   ports.MasterCurrenciesRepository
	repoMasterFAD  ports.MasterFacilitiesDescriptionRepository
	repoMasterFA   ports.MasterFacilitiesRepository
	repoMasterFAGD ports.MasterFacilityGroupsDescriptionRepository
	repoMasterFAG  ports.MasterFacilityGroupsRepository
	repoMasterFAT  ports.MasterFacilityTypologiesRepository
	repoMasterGCAD ports.MasterGroupCategoriesDescriptionRepository
	repoMasterGCA  ports.MasterGroupCategoriesRepository
	repoMasterHT   ports.MasterHotelTypeRepository
	repoMasterITD  ports.MasterImageTypeDescriptionRepository
	repoMasterIT   ports.MasterImageTypesRepository
	repoMasterISD  ports.MasterIssuesDescriptionRepository
	repoMasterIS   ports.MasterIssuesRepository
	repoMasterLD   ports.MasterLanguagesDescriptionRepository
	repoMasterL    ports.MasterLanguagesRepository
	repoMasterPRD  ports.MasterPromotionsDescriptionRepository
	repoMasterPRN  ports.MasterPromotionsNameRepository
	repoMasterPR   ports.MasterPromotionsRepository
	repoMasterROBT ports.MasterRoomBedsTypeRepository
	repoMasterROCD ports.MasterRoomsCharacteristicDescriptionRepository
	repoMasterRO   ports.MasterRoomsRepository
	repoMasterROTD ports.MasterRoomsTypeDescriptionRepository
	repoMasterSED  ports.MasterSegmentsDescriptionRepository
	repoMasterSE   ports.MasterSegmentsRepository
	repoMasterSU   ports.MasterSurroundingsRepository
	repoMasterTED  ports.MasterTerminalsDescriptionRepository
	repoMasterTEN  ports.MasterTerminalsNameRepository
	repoMasterTE   ports.MasterTerminalsRepository
}

func NewService(db *gorm.DB) *service {
	repoMasterACD := masterrepo.NewAccommodationsDescriptionRepo(db)
	repoMasterAC := masterrepo.NewAccommodationsRepo(db)
	repoMasterAM := masterrepo.NewAmenitiesRepo(db)
	repoMasterBOD := masterrepo.NewBoardsDescriptionRepo(db)
	repoMasterBO := masterrepo.NewBoardsRepo(db)
	repoMasterCAD := masterrepo.NewCategoriesDescriptionRepo(db)
	repoMasterCA := masterrepo.NewCategoriesRepo(db)
	repoMasterCHD := masterrepo.NewChainsDescriptionRepo(db)
	repoMasterCH := masterrepo.NewChainsRepo(db)
	repoMasterCUD := masterrepo.NewCurrenciesDescriptionRepo(db)
	repoMasterCU := masterrepo.NewCurrenciesRepo(db)
	repoMasterFAD := masterrepo.NewFacilitiesDescriptionRepo(db)
	repoMasterFA := masterrepo.NewFacilitiesRepo(db)
	repoMasterFAGD := masterrepo.NewFacilityGroupsDescriptionRepo(db)
	repoMasterFAG := masterrepo.NewFacilityGroupsRepo(db)
	repoMasterFAT := masterrepo.NewFacilityTypologiesRepo(db)
	repoMasterGCAD := masterrepo.NewGroupCategoriesDescriptionRepo(db)
	repoMasterGCA := masterrepo.NewGroupCategoriesRepo(db)
	repoMasterHT := masterrepo.NewHotelTypeRepo(db)
	repoMasterITD := masterrepo.NewImageTypeDescriptionRepo(db)
	repoMasterIT := masterrepo.NewImageTypesRepo(db)
	repoMasterISD := masterrepo.NewIssuesDescriptionRepo(db)
	repoMasterIS := masterrepo.NewIssuesRepo(db)
	repoMasterLD := masterrepo.NewLanguagesDescriptionRepo(db)
	repoMasterL := masterrepo.NewLanguagesRepo(db)
	repoMasterPRD := masterrepo.NewPromotionsDescriptionRepo(db)
	repoMasterPRN := masterrepo.NewPromotionsNameRepo(db)
	repoMasterPR := masterrepo.NewPromotionsRepo(db)
	repoMasterROBT := masterrepo.NewRoomBedsTypeRepo(db)
	repoMasterROCD := masterrepo.NewRoomsCharacteristicDescriptionRepo(db)
	repoMasterRO := masterrepo.NewRoomsRepo(db)
	repoMasterROTD := masterrepo.NewRoomsTypeDescriptionRepo(db)
	repoMasterSED := masterrepo.NewSegmentsDescriptionRepo(db)
	repoMasterSE := masterrepo.NewSegmentsRepo(db)
	repoMasterSU := masterrepo.NewSurroundingsRepo(db)
	repoMasterTED := masterrepo.NewTerminalsDescriptionRepo(db)
	repoMasterTEN := masterrepo.NewTerminalsNameRepo(db)
	repoMasterTE := masterrepo.NewTerminalsRepo(db)

	return &service{
		repoMasterACD:  repoMasterACD,
		repoMasterAC:   repoMasterAC,
		repoMasterAM:   repoMasterAM,
		repoMasterBOD:  repoMasterBOD,
		repoMasterBO:   repoMasterBO,
		repoMasterCAD:  repoMasterCAD,
		repoMasterCA:   repoMasterCA,
		repoMasterCHD:  repoMasterCHD,
		repoMasterCH:   repoMasterCH,
		repoMasterCUD:  repoMasterCUD,
		repoMasterCU:   repoMasterCU,
		repoMasterFAD:  repoMasterFAD,
		repoMasterFA:   repoMasterFA,
		repoMasterFAGD: repoMasterFAGD,
		repoMasterFAG:  repoMasterFAG,
		repoMasterFAT:  repoMasterFAT,
		repoMasterGCAD: repoMasterGCAD,
		repoMasterGCA:  repoMasterGCA,
		repoMasterHT:   repoMasterHT,
		repoMasterITD:  repoMasterITD,
		repoMasterIT:   repoMasterIT,
		repoMasterISD:  repoMasterISD,
		repoMasterIS:   repoMasterIS,
		repoMasterLD:   repoMasterLD,
		repoMasterL:    repoMasterL,
		repoMasterPRD:  repoMasterPRD,
		repoMasterPRN:  repoMasterPRN,
		repoMasterPR:   repoMasterPR,
		repoMasterROBT: repoMasterROBT,
		repoMasterROCD: repoMasterROCD,
		repoMasterRO:   repoMasterRO,
		repoMasterROTD: repoMasterROTD,
		repoMasterSED:  repoMasterSED,
		repoMasterSE:   repoMasterSE,
		repoMasterSU:   repoMasterSU,
		repoMasterTED:  repoMasterTED,
		repoMasterTEN:  repoMasterTEN,
		repoMasterTE:   repoMasterTE,
	}
}
