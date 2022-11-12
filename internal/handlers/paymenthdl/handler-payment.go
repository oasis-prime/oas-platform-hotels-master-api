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
	transactionId, err := strconv.ParseInt(*input.ClientReference, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%s", "ParseIntClientReference")
	}

	booking, err := h.servBooking.GetPayLinkId(fmt.Sprintf("SM-%d", transactionId))

	if err == nil && booking != nil {
		copier.Copy(&display, &booking)
		display.Reference = &booking.ClientReference

		return &display, nil
	}

	getChillPay, err := h.servPaymentService.GetDetailByTransctionID(&chillpaydm.TransctionIDRequest{TransactionId: fmt.Sprintf("%d", transactionId)})

	if err != nil {
		return nil, fmt.Errorf("%s", "ClientReferenceNotFound")
	}

	paymentStatus := getChillPay.Data.PaymentStatus
	payLinkId := getChillPay.Data.PayLinkID
	if paymentStatus == "Actived" || paymentStatus == "Waiting" {
		return nil, fmt.Errorf("%s", "PaymentGatewayIsClose")
	}

	if paymentStatus == "Expired" ||
		paymentStatus == "Deleted" ||
		paymentStatus == "Cancelled" ||
		paymentStatus == "Closed" {
		return nil, fmt.Errorf("%s", "ClientReferenceIsExpired")
	}

	payment, err := h.servPaymentService.GetPayment(payLinkId)
	if err != nil || payment == nil {
		return nil, fmt.Errorf("%s", "ClientReferenceNotFound")
	}

	if payment.Status != paymentStatus {
		payment.Status = getChillPay.Data.PaymentStatus
		copier.Copy(payment, &getChillPay.Data)

		layout := "02/01/2006 15:04:05"
		transactionDate, _ := time.Parse(layout, getChillPay.Data.TransactionDate)
		paymentDate, _ := time.Parse(layout, getChillPay.Data.PaymentDate)

		payment.TransactionDate = &transactionDate
		payment.PaymentDate = &paymentDate

		fmt.Println(getChillPay.Data)
		fmt.Println(payment)

		_, err = h.servPaymentService.UpdatePayment(payment.Code, payment)
		if err != nil {
			return nil, fmt.Errorf("%s", "UpdatePaymentError")
		}
	}

	if payment.Status == "Success" {
		fmt.Println("running Success")
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

		fmt.Println("running Success!!")

		mb := &hotelbedsdm.BookingsRequest{
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
		}

		fmt.Printf("running Success: %v \n", mb)
		booking, err := h.servHotelbeds.Booking(mb)

		if err != nil {
			return nil, err
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

		hotelImage, _, err := h.servHotel.GetHotelImages(hoteldm.GetAllHotelImagesRequest{
			HotelCode: &hotelcode,
			HotelType: &hoteltype,
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				IsOffset: true,
				Offset:   0,
				Limit:    1,
				Language: (*langenums.Language)(&input.Language),
			},
		})

		if err != nil || len(hotelImage) < 1 {
			return nil, fmt.Errorf("%s", "HotelImageNotFound")
		}

		hotelCoordinates, _, err := h.servHotel.GetCoordinates(hoteldm.GetAllHotelCoordinatesRequest{
			HotelCode: &hotelcode,
			HotelType: &hoteltype,
			GetAllRequestBasic: hoteldm.GetAllRequestBasic{
				IsOffset: true,
				Offset:   0,
				Limit:    1,
				Language: (*langenums.Language)(&input.Language),
			},
		})

		if err != nil || len(hotelCoordinates) < 1 {
			return nil, fmt.Errorf("%s", "HotelCoordinatesNotFound")
		}

		checkIn, err := time.Parse(timeFormat, booking.Booking.Hotel.CheckIn)
		if err != nil {
			return nil, fmt.Errorf("%s", "")
		}
		checkOut, err := time.Parse(timeFormat, booking.Booking.Hotel.CheckOut)
		if err != nil {
			return nil, fmt.Errorf("%s", "")
		}
		days := checkOut.Sub(checkIn)

		if booking != nil {
			h.servPaymentService.BookingMail(&domain.PublisherBookingEmail{
				Adults:          fmt.Sprintf("%d", maxAdults),
				BookingID:       booking.Booking.ClientReference,
				CategoryName:    booking.Booking.Hotel.CategoryName,
				CheckIn:         booking.Booking.Hotel.CheckIn,
				CheckOut:        booking.Booking.Hotel.CheckOut,
				Cost:            fmt.Sprintf("%.2f", payment.Amount),
				Days:            fmt.Sprintf("%.f", days.Hours()/24),
				DestinationName: booking.Booking.Hotel.DestinationName,
				Email:           payment.Email,
				HotelAddress:    hoteladdress[0].Content,
				HotelName:       booking.Booking.Hotel.Name,
				HotelImage:      hotelImage[0].Path,
				Latitude:        fmt.Sprintf("%f", hotelCoordinates[0].Latitude),
				Logo:            "",
				Longitude:       fmt.Sprintf("%f", hotelCoordinates[0].Longitude),
				PostalCode:      hotel.PostalCode,
				RoomAmount:      fmt.Sprintf("%d", maxRoom),
				RoomType:        booking.Booking.Hotel.Rooms[0].Name,
				ZoneName:        booking.Booking.Hotel.ZoneName,
			})

			copier.Copy(&display, &booking.Booking)
			display.Reference = &booking.Booking.ClientReference

			return &display, nil
		}
	}

	return &display, fmt.Errorf("%v", "ErrorBooking")
}
