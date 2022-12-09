package tickerserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/customerdm"
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	tickerRepo         ports.TickerRepository
	customerTickerRepo ports.CustomerTickerRepository
}

func NewService(tickerRepo ports.TickerRepository, customerTickerRepo ports.CustomerTickerRepository) *service {
	return &service{
		tickerRepo:         tickerRepo,
		customerTickerRepo: customerTickerRepo,
	}
}

func (svc *service) TickerGetAll(condition masterdm.GetAllRequestBasic) (results []*masterrepo.Ticker, totalRows int64, err error) {
	return svc.tickerRepo.GetAll(condition)
}

func (svc *service) TickerGet(argID uint) (record *masterrepo.Ticker, err error) {
	return svc.tickerRepo.Get(argID)
}

func (svc *service) CustomerTickerGetAll(condition customerdm.GetAllRequestBasic) (results []*customerrepo.CustomerTicker, totalRows int64, err error) {
	return svc.customerTickerRepo.GetAll(condition)
}

func (svc *service) CustomerTickerGet(argID uint) (record *customerrepo.CustomerTicker, err error) {
	return svc.customerTickerRepo.Get(argID)
}
