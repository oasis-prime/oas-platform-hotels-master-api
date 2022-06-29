package ports

import "github.com/oasis-prime/oas-platform-hotels-master-api/repositories"

type MemberRepository interface {
	GetAll(page, pagesize int, order string) (results []*repositories.SystemsMember, totalRows int64, err error)
}
