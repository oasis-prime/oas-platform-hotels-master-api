package paymenthdl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servHotelbeds      ports.HotelbedsService
	servPaymentService ports.PaymentService
}

func NewHandler(
	servHotelbeds ports.HotelbedsService,
	servPaymentService ports.PaymentService,
) *Handler {
	return &Handler{
		servHotelbeds:      servHotelbeds,
		servPaymentService: servPaymentService,
	}
}

// Payment(ctx context.Context, input model.PaymentInput) (*model.PaymentData, error)
func (h *Handler) Payment(ctx context.Context, input model.PaymentInput) (display *model.PaymentData, err error) {

	rate, err := h.servHotelbeds.CheckRate(&hotelbedsdm.CheckRatesRequest{Rooms: []hotelbedsdm.RateKey{{RateKey: input.RateKey}}})
	if err != nil {
		return nil, err
	}

	startDate := time.Now()
	expiredDate := time.Now().Add(time.Minute * 5)

	pay, err := h.servPaymentService.Generate(&chillpaydm.PaylinkGenerateRequest{
		ProductImage:       "",
		ProductName:        rate.Hotel.Name,
		ProductDescription: rate.Hotel.ZoneName,
		PaymentLimit:       "",
		StartDate:          startDate.Format("02/01/2006 15:04:05"),
		ExpiredDate:        expiredDate.Format("02/01/2006 15:04:05"),
		Currency:           "THB",
		Amount:             strings.Replace(rate.Hotel.TotalNet, ".", "", 3),
	})

	if err != nil {
		return nil, err
	}

	_, err = h.servPaymentService.CreatePayment(&customerrepo.CustomerPayment{
		Email:       input.Email,
		Name:        input.Name,
		Surname:     input.Surname,
		PhoneNumber: input.PhoneNumber,
		RateKey:     input.RateKey,
		PayLinkId:   pay.Data.PayLinkID,
		Status:      pay.Data.Status,
	})

	if err != nil {
		return nil, err
	}

	display = &model.PaymentData{
		OrderNumber: fmt.Sprintf("%d", pay.Data.PayLinkID),
		PaymentURL:  pay.Data.PaymentURL,
		QRImage:     pay.Data.PaymentURL,
	}

	return display, nil
}
