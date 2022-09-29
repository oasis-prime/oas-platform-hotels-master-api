package paymentserv

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type service struct {
	cillpayHttp ports.PaymentChillpayHttp
	repo        ports.PaymentRepository
	pub         *pubsub.Client
}

func NewService(cillpayHttp ports.PaymentChillpayHttp, repo ports.PaymentRepository, pub *pubsub.Client) *service {
	return &service{
		cillpayHttp: cillpayHttp,
		repo:        repo,
		pub:         pub,
	}
}

func (svc *service) Generate(condition *chillpaydm.PaylinkGenerateRequest) (response *chillpaydm.PaylinkGenerateResponse, err error) {
	response, err = svc.cillpayHttp.GetPaylinkGenerate(condition)
	return response, err
}

func (svc *service) GetChillPay(argCode uint) (response *chillpaydm.PaylinkGenerateResponse, err error) {
	response, err = svc.cillpayHttp.GetPaylinkDetail(&chillpaydm.PaylinkDetailsRequest{
		PayLinkId: fmt.Sprintf("%d", argCode),
	})
	return response, err
}

func (svc *service) CreatePayment(record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, err error) {
	result, _, err = svc.repo.Create(record)
	return result, err
}

func (svc *service) UpdatePayment(argCode uint, record *customerrepo.CustomerPayment) (result *customerrepo.CustomerPayment, err error) {
	result, _, err = svc.repo.Update(argCode, record)
	return result, err
}

func (svc *service) BookingMail(condition *domain.PublisherBookingEmail) (err error) {
	topic := svc.pub.Topic("oas-trigger-email-booking")

	byteData, _ := json.Marshal(condition)

	result := topic.Publish(context.TODO(), &pubsub.Message{
		Data: []byte(byteData),
		Attributes: map[string]string{
			"adults":          condition.Adults,
			"bookingID":       condition.BookingID,
			"categoryName":    condition.CategoryName,
			"checkIn":         condition.CheckIn,
			"checkOut":        condition.CheckOut,
			"cost":            condition.Cost,
			"days":            condition.Days,
			"destinationName": condition.DestinationName,
			"email":           condition.Email,
			"hotelAddress":    condition.HotelAddress,
			"hotelName":       condition.HotelName,
			"logo":            condition.Logo,
			"postalCode":      condition.PostalCode,
			"roomAmount":      condition.RoomAmount,
			"roomType":        condition.RoomType,
			"latitude":        condition.Latitude,
			"longitude":       condition.Longitude,
			"hotelImage":      condition.HotelImage,
		},
	})

	id, err := result.Get(context.Background())
	if err != nil {
		return err
	}

	log.Printf("Public message to pub/sub complate : %s", id)

	return err
}

func (svc *service) GetPayment(payLinkId int) (result *customerrepo.CustomerPayment, err error) {
	result, err = svc.repo.GetPayLinkId(payLinkId)
	return result, err
}

func (svc *service) GetDetailByTransctionID(condition *chillpaydm.TransctionIDRequest) (response *chillpaydm.DetailByTransctionResponse, err error) {
	return svc.cillpayHttp.GetDetailByTransctionID(condition)
}
