package hotelbedssvc

import (
	"encoding/json"
	"fmt"

	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
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

	f, _ := json.Marshal(response)
	fmt.Println(string(f))

	return nil
}
