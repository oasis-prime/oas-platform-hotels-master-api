package ports

import (
	"github.com/oasis-prime/oas-platform-core/repositories"
)

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*repositories.SystemsMember, totalRows int64, err error)
}

type HotelbedsContentRepository interface {
	GetAll(condition repositories.GetAllRequestBasic) (results []*repositories.HotelsRepo, totalRows int64, err error)
	Get(argID uint) (record *repositories.Hotels, err error)
	CreateBatch(record *[]repositories.Hotels) (RowsAffected int64, err error)
	Create(record *repositories.Hotels) (result *repositories.Hotels, RowsAffected int64, err error)
	Update(argID uint, updated *repositories.Hotels) (result *repositories.Hotels, RowsAffected int64, err error)
	Delete(argID uint32) (rowsAffected int64, err error)
}
