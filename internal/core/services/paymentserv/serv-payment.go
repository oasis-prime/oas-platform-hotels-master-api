package paymentserv

import (
	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	cillpayHttp ports.PaymentChillpayHttp
	repo        ports.PaymentRepository
}

func NewService(cillpayHttp ports.PaymentChillpayHttp, repo ports.PaymentRepository) *service {
	return &service{
		cillpayHttp: cillpayHttp,
		repo:        repo,
	}
}

func (svc *service) Generate(condition *chillpaydm.PaylinkGenerateRequest) (response *chillpaydm.PaylinkGenerateResponse, err error) {
	response, err = svc.cillpayHttp.GetPaylinkGenerate(condition)
	return response, err
}

func (svc *service) CreatePayment(record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, err error) {
	result, _, err = svc.repo.Create(record)
	return result, err
}
