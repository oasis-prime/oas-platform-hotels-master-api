package paymenthdl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"github.com/oasis-prime/oas-platform-core/domain/chillpaydm"
	"github.com/oasis-prime/oas-platform-core/domain/hotelbedsdm"
	"github.com/oasis-prime/oas-platform-core/domain/hoteldm"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/htenums"
	"github.com/oasis-prime/oas-platform-core/repositories/enums/langenums"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/model"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/domain"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/ports"
)

type Handler struct {
	servHotelbeds      ports.HotelbedsService
	servPaymentService ports.PaymentService
	servHotel          ports.HotelsService
	servBooking        ports.BookingService
}

func NewHandler(
	servHotelbeds ports.HotelbedsService,
	servPaymentService ports.PaymentService,
	servHotel ports.HotelsService,
	servBooking ports.BookingService,
) *Handler {
	return &Handler{
		servHotelbeds:      servHotelbeds,
		servPaymentService: servPaymentService,
		servHotel:          servHotel,
		servBooking:        servBooking,
	}
}

const (
	timeFormat = "2006-01-02"
)

func (h *Handler) GetPayment(ctx context.Context, input model.GetPaymentInput) (display *model.PaymentData, err error) {
	argCode, err := strconv.ParseInt(input.OrderNumber, 10, 32)
	if err != nil {
		return nil, err
	}

	_, err = h.servPaymentService.GetPayment(int(argCode))
	if err != nil {
		return nil, err
	}

	response, err := h.servPaymentService.GetChillPay(uint(argCode))
	if err != nil {
		return nil, err
	}

	f, _ := json.Marshal(response)
	fmt.Println(string(f))

	return display, nil
}

// Payment(ctx context.Context, input model.PaymentInput) (*model.PaymentData, error)
func (h *Handler) Payment(ctx context.Context, input model.PaymentInput) (display *model.PaymentData, err error) {

	rate, err := h.servHotelbeds.CheckRate(&hotelbedsdm.CheckRatesRequest{
		Rooms: []hotelbedsdm.RateKey{
			{
				RateKey: input.RateKey,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(rate.Hotel.TotalNet, 10)
	if err != nil {
		return nil, err
	}

	startDate := time.Now()
	expiredDate := time.Now().Add(time.Minute * 5)

	condition := &chillpaydm.PaylinkGenerateRequest{
		ProductImage:       "",
		ProductName:        rate.Hotel.Name,
		ProductDescription: rate.Hotel.ZoneName,
		PaymentLimit:       "",
		StartDate:          startDate.Format("02/01/2006 15:04:05"),
		ExpiredDate:        expiredDate.Format("02/01/2006 15:04:05"),
		Currency:           "THB",
		Amount:             strings.Replace(rate.Hotel.TotalNet, ".", "", 3),
	}

	pay, err := h.servPaymentService.Generate(condition)

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
		Amount:      amount,
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

func (h *Handler) Booking(ctx context.Context, input model.BookingInput) (*model.BookingData, error) {
	display := model.BookingData{}

	if input.ClientReference == nil {
		return nil, fmt.Errorf("%s", "RequiredFieldClientReference")
	}
	payLinkId, err := strconv.ParseInt(*input.ClientReference, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s", "ParseIntClientReference")
	}

	booking, _ := h.servBooking.GetPayLinkId(fmt.Sprintf("SM-%d", payLinkId))

	if booking != nil {
		copier.Copy(&display, &booking)
		display.Reference = &booking.ClientReference

		return &display, nil
	}

	payment, err := h.servPaymentService.GetPayment(int(payLinkId))
	if err != nil || payment == nil {
		return nil, fmt.Errorf("%s", "ClientReferenceNotFound")
	}

	if payment.Status == "Expired" ||
		payment.Status == "Deleted" ||
		payment.Status == "Cancelled" ||
		payment.Status == "Closed" {
		return nil, fmt.Errorf("%s", "ClientReferenceIsExpired")
	}

	if payment.Status == "Actived" || payment.Status == "Waiting" {
		getChillPay, err := h.servPaymentService.GetChillPay(uint(payLinkId))
		if err != nil || getChillPay == nil {
			return nil, fmt.Errorf("%s", "PaymentGatewayIsClose")
		}

		payment.Status = getChillPay.Data.Status

		_, err = h.servPaymentService.UpdatePayment(payment.Code, payment)
		if err != nil {
			return nil, fmt.Errorf("%s", "UpdatePaymentError")
		}
	}

	if payment.Status == "Success" {
		checkRate, err := h.servHotelbeds.CheckRate(&hotelbedsdm.CheckRatesRequest{
			Rooms: []hotelbedsdm.RateKey{
				{RateKey: payment.RateKey},
			},
		})

		if err != nil {
			return nil, fmt.Errorf("%s", "CheckRateError")
		}

		paxes := []hotelbedsdm.BookingPaxes{}
		var maxRoom, maxAdults int
		if len(checkRate.Hotel.Rooms) > 0 && len(checkRate.Hotel.Rooms[0].Rates) > 0 {
			maxRoom = checkRate.Hotel.Rooms[0].Rates[0].Rooms
			maxAdults = checkRate.Hotel.Rooms[0].Rates[0].Adults

			for r := 0; r < maxRoom; r++ {
				for a := 0; a < maxAdults; a++ {
					paxes = append(paxes, hotelbedsdm.BookingPaxes{
						RoomID:  r + 1,
						Type:    "AD",
						Name:    payment.Name,
						Surname: payment.Surname,
					})
				}
			}
		}

		booking, err := h.servHotelbeds.Booking(&hotelbedsdm.BookingsRequest{
			Language: hotelbedsdm.Language(input.Language),
			Holder: hotelbedsdm.BookingHolder{
				Name:    payment.Name,
				Surname: payment.Surname,
			},
			Rooms: []hotelbedsdm.BookingRooms{
				{
					RateKey: payment.RateKey,
					Paxes:   paxes,
				},
			},
			ClientReference: fmt.Sprintf("SM-%d", payment.PayLinkId),
			Remark:          "Smartgo.life",
			Tolerance:       2,
		})

		if err != nil {
			return nil, fmt.Errorf("%s", "BookingError")
		}

		fmt.Printf("booking: %v \n", booking)

		createBooking := &customerrepo.CustomerBooking{}
		copier.CopyWithOption(createBooking, &booking.Booking, copier.Option{IgnoreEmpty: true})
		fmt.Printf("createBooking: %v \n", createBooking)

		_, _, err = h.servBooking.Create(createBooking)
		if err != nil {
			return nil, fmt.Errorf("%s", "SaveBookingError")
		}

		hotelcode := uint(booking.Booking.Hotel.Code)
		hoteltype := htenums.Hotelbeds
		hotel, err := h.servHotel.GetHotel(model.HotelInput{
			Code:     int(hotelcode),
			Language: input.Language,
		})
		if err != nil || hotel == nil {
			return nil, fmt.Errorf("%s", "HotelNotFound")
		}
		hoteladdress, _, err := h.servHotel.GetHotelAddress(hoteldm.GetAllHotelAddressRequest{
			HotelCode: &hotelcode,
			HotelType: &hoteltype,
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				IsOffset: true,
				Offset:   0,
				Limit:    1,
				Language: (*langenums.Language)(&input.Language),
			},
		})

		if err != nil || len(hoteladdress) < 1 {
			return nil, fmt.Errorf("%s", "HotelAddressNotFound")
		}

		checkIn, err := time.Parse(timeFormat, booking.Booking.Hotel.CheckIn)
		if err != nil {
			return nil, fmt.Errorf("%s", "")
		}
		checkOut, err := time.Parse(timeFormat, booking.Booking.Hotel.CheckOut)
		if err != nil {
			return nil, fmt.Errorf("%s", "")
		}
		days := checkIn.Sub(checkOut)

		if booking != nil {
			h.servPaymentService.BookingMail(&domain.PublisherBookingEmail{
				Logo:            "Logo",
				BookingID:       booking.Booking.ClientReference,
				HotelName:       booking.Booking.Hotel.Name,
				CategoryName:    booking.Booking.Hotel.CategoryName,
				HotelAddress:    hoteladdress[0].Content,
				PostalCode:      hotel.PostalCode,
				ZoneName:        booking.Booking.Hotel.ZoneName,
				DestinationName: booking.Booking.Hotel.DestinationName,
				Days:            fmt.Sprintf("%.f", days.Hours()/24),
				CheckIn:         booking.Booking.Hotel.CheckIn,
				CheckOut:        booking.Booking.Hotel.CheckOut,
				RoomType:        booking.Booking.Hotel.Rooms[0].Name,
				RoomAmount:      fmt.Sprintf("%d", maxRoom),
				Adults:          fmt.Sprintf("%d", maxAdults),
				Cost:            fmt.Sprintf("%d", payment.Amount),
				Email:           payment.Email,
			})

			copier.Copy(&display, &booking.Booking)
			display.Reference = &booking.Booking.ClientReference

			return &display, nil
		}
	}

	return &display, fmt.Errorf("%v", "ErrorBooking")
}
