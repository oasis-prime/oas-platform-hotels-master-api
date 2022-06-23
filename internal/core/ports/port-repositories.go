package ports

import "github.com/oasis-prime/oas-platform-hotels-master-api/internal/repositories"

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*repositories.Member, totalRows int64, err error)
}
