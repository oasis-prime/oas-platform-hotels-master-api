package bookingserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/customerdm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	bookingHotelRepo ports.CustomerBookingRepository
}

func NewService(
	bookingHotelRepo ports.CustomerBookingRepository,
) *service {
	return &service{
		bookingHotelRepo: bookingHotelRepo,
	}
}

func (svc *service) GetAll(condition customerdm.GetAllRequestBasic) (results []*customerrepo.CustomerBooking, totalRows int64, err error) {
	return svc.bookingHotelRepo.GetAll(condition)
}

func (svc *service) Get(argID uint) (record *customerrepo.CustomerBooking, err error) {
	return svc.bookingHotelRepo.Get(argID)
}

func (svc *service) Create(record *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error) {
	return svc.bookingHotelRepo.Create(record)
}

func (svc *service) Update(argID uint, updated *customerrepo.CustomerBooking) (result *customerrepo.CustomerBooking, RowsAffected int64, err error) {
	return svc.bookingHotelRepo.Update(argID, updated)
}

func (svc *service) GetPayLinkId(payLinkId string) (record *customerrepo.CustomerBooking, err error) {
	return svc.bookingHotelRepo.GetPayLinkId(payLinkId)
}
