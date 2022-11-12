package popularserv

import (
	"cloud.google.com/go/pubsub"
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	repo ports.PopularRepository
	pub  *pubsub.Client
}

func NewService(repo ports.PopularRepository) *service {
	return &service{
		repo: repo,
	}
}

func (svc *service) GetAll(condition masterdm.GetCustomerPopularRequest) (results []*masterrepo.Popular, totalRows int64, err error) {
	return svc.repo.GetAll(condition)
}

func (svc *service) Get(condition masterdm.GetCustomerPopularRequest) (record *masterrepo.Popular, err error) {
	return svc.repo.Get(condition)
}
