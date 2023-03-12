package protocol

import (
	"context"
	"fmt"

	"github.com/oasis-prime/oas-platform-core/http/chillpayhttp"
	"github.com/oasis-prime/oas-platform-core/http/hotelbedshttp"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/masterrepo"
	"github.com/oasis-prime/oas-platform-firebase-core/repositories/firebaserepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/bookingserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/googleserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelbedsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/masterserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/memberserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/paymentserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/popularserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/subscriptionserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/tickerserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/googlehdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/memberhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/paymenthdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/popularhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/subscriptionhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/tickerhdl"
)

func paymentHandlerInit() *paymenthdl.Handler {
	apprequestChillpay := chillpayhttp.NewRequester(con.Chillpay.Merchantcode, con.Chillpay.Apikey)
	customerPaymentRepo := customerrepo.NewCustomerPaymentRepo(db)
	cillpayHttp := chillpayhttp.NewChillpayHTTP(con.Chillpay.Url, con.Chillpay.MD5, apprequestChillpay)
	servPayment := paymentserv.NewService(cillpayHttp, customerPaymentRepo, pub)

	repoHotels := hotelrepo.NewHotelsRepo(db)
	repoHotelName := hotelrepo.NewHotelNameRepo(db)
	repoHotelDescription := hotelrepo.NewHotelDescriptionRepo(db)
	repoHotelInterestPoints := hotelrepo.NewHotelInterestPointsRepo(db)
	repoHotelIssues := hotelrepo.NewHotelIssuesRepo(db)
	repoHotelFacility := hotelrepo.NewHotelFacilityRepo(db)
	repoHotelRooms := hotelrepo.NewHotelRoomsRepo(db)
	repoHotelRoomsFacilities := hotelrepo.NewHotelRoomFacilitiesRepo(db)
	repoHotelRoomsStay := hotelrepo.NewHotelRoomStayRepo(db)
	repoHotelPhones := hotelrepo.NewHotelPhoneRepo(db)
	repoHotelCity := hotelrepo.NewHotelCityRepo(db)
	repoHotelAddress := hotelrepo.NewHotelAddressRepo(db)
	repoHotelImages := hotelrepo.NewHotelImagesRepo(db)
	repoCoordinates := hotelrepo.NewHotelCoordinatesRepo(db)

	repoBooking := customerrepo.NewCustomerBookingRepo(db)

	apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
	hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
	hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)
	hotelsServ := hotelsserv.NewService(
		repoHotels,
		repoHotelName,
		repoHotelDescription,
		repoHotelInterestPoints,
		repoHotelIssues,
		repoHotelFacility,
		repoHotelRooms,
		repoHotelRoomsFacilities,
		repoHotelRoomsStay,
		repoHotelPhones,
		repoHotelCity,
		repoHotelAddress,
		repoHotelImages,
		repoCoordinates,
	)
	bookingServ := bookingserv.NewService(repoBooking)
	return paymenthdl.NewHandler(hotelbedsServ, servPayment, hotelsServ, bookingServ)
}

func googleHandlerInit() *googlehdl.Handler {
	servPlance := googleserv.NewService("AIzaSyDOS6MedhWdvMMyRvGkRTROvTXnf8exZW8")
	return googlehdl.NewHandler(servPlance)
}

func hotelsHandlerInit() *hotelshdl.Handler {
	repoHotels := hotelrepo.NewHotelsRepo(db)
	repoHotelName := hotelrepo.NewHotelNameRepo(db)
	repoHotelDescription := hotelrepo.NewHotelDescriptionRepo(db)
	repoHotelInterestPoints := hotelrepo.NewHotelInterestPointsRepo(db)
	repoHotelIssues := hotelrepo.NewHotelIssuesRepo(db)
	repoHotelFacility := hotelrepo.NewHotelFacilityRepo(db)
	repoHotelRooms := hotelrepo.NewHotelRoomsRepo(db)
	repoHotelRoomsFacilities := hotelrepo.NewHotelRoomFacilitiesRepo(db)
	repoHotelRoomsStay := hotelrepo.NewHotelRoomStayRepo(db)
	repoHotelPhones := hotelrepo.NewHotelPhoneRepo(db)
	repoHotelCity := hotelrepo.NewHotelCityRepo(db)
	repoHotelAddress := hotelrepo.NewHotelAddressRepo(db)
	repoHotelImages := hotelrepo.NewHotelImagesRepo(db)
	repoCoordinates := hotelrepo.NewHotelCoordinatesRepo(db)

	hotelsServ := hotelsserv.NewService(
		repoHotels,
		repoHotelName,
		repoHotelDescription,
		repoHotelInterestPoints,
		repoHotelIssues,
		repoHotelFacility,
		repoHotelRooms,
		repoHotelRoomsFacilities,
		repoHotelRoomsStay,
		repoHotelPhones,
		repoHotelCity,
		repoHotelAddress,
		repoHotelImages,
		repoCoordinates,
	)

	masterServ := masterserv.NewService(db)

	return hotelshdl.NewHandler(hotelsServ, masterServ)
}

func hotelbedsHandlerInit() *hotelbedshdl.Handler {
	apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
	hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
	hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)
	return hotelbedshdl.NewHandler(hotelbedsServ)
}

func memberHandlerInit() *memberhdl.Handler {
	memberRepo := customerrepo.NewMemberRepo(db)
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(fmt.Errorf("client: %v", err))
	}

	firebaseMemberRepo := firebaserepo.NewFirebaseMemberRepo(app, client)
	memberServ := memberserv.NewService(memberRepo, pub, firebaseMemberRepo)
	return memberhdl.NewHandler(memberServ)
}

func popularHandlerInit() *popularhdl.Handler {
	repo := masterrepo.NewPopularRepo(db)
	popularServ := popularserv.NewService(repo)

	return popularhdl.NewHandler(popularServ)
}

func tickerHandlerInit() *tickerhdl.Handler {
	tickerRepo := masterrepo.NewTickerRepo(db)
	customerTickerRepo := customerrepo.NewCustomerTickergRepo(db)

	tickerServ := tickerserv.NewService(tickerRepo, customerTickerRepo)

	return tickerhdl.NewHandler(tickerServ)
}

func subscriptionHandlerInit() *subscriptionhdl.Handler {
	repo := customerrepo.NewCustomerSubscriptionRepo(db)
	serv := subscriptionserv.NewService(repo)

	return subscriptionhdl.NewHandler(serv)
}
