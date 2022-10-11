package masterserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/masterdm"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
)

func (svc *service) GetRoomsStay(
	condition masterdm.GetCustomerPopularRequest,
) (results []*masterrepo.Popular, totalRows int64, err error) {
	return svc.Popular.GetAll(condition)
}
