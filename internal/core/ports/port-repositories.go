package ports

import (
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
)

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*customerrepo.CustomerMember, totalRows int64, err error)
}

type HotelsRepository interface {
	GetAll(condition hoteldm.GetAllRequestBasic) (results []*hotelrepo.Hotels, totalRows int64, err error)
	Get(argID uint) (record *hotelrepo.Hotels, err error)
	CreateBatch(record *[]hotelrepo.Hotels) (RowsAffected int64, err error)
	Create(record *hotelrepo.Hotels) (result *hotelrepo.Hotels, RowsAffected int64, err error)
	Update(argID uint, updated *hotelrepo.Hotels) (result *hotelrepo.Hotels, RowsAffected int64, err error)
	Delete(argID uint32) (rowsAffected int64, err error)
	GetByCoordinates(condition hoteldm.GetByCoordinatesRequest) (record []*hotelrepo.Hotels, err error)
}

type HotelbedsBookingRepository interface {
}
