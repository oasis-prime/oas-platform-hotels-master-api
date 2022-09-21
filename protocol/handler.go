package protocol

import (
	"github.com/oasis-prime/oas-platform-core/http/chillpayhttp"
	"github.com/oasis-prime/oas-platform-core/http/hotelbedshttp"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/googleserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelbedsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/memberserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/paymentserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/googlehdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/memberhdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/paymenthdl"
)

func paymentHandlerInit() *paymenthdl.Handler {
	apprequestChillpay := chillpayhttp.NewRequester(con.Chillpay.Merchantcode, con.Chillpay.Apikey)
	customerPaymentRepo := customerrepo.NewCustomerPaymentRepo(db)
	cillpayHttp := chillpayhttp.NewChillpayHTTP(con.Chillpay.Url, con.Chillpay.MD5, apprequestChillpay)
	servPayment := paymentserv.NewService(cillpayHttp, customerPaymentRepo)

	apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
	hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
	hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)

	return paymenthdl.NewHandler(hotelbedsServ, servPayment)
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
	)
	return hotelshdl.NewHandler(hotelsServ)
}

func hotelbedsHandlerInit() *hotelbedshdl.Handler {
	apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
	hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
	hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)
	return hotelbedshdl.NewHandler(hotelbedsServ)
}

func memberHandlerInit() *memberhdl.Handler {
	memberRepo := customerrepo.NewMemberRepo(db)
	memberServ := memberserv.NewService(memberRepo, con.Google.Projectid, con.Google.Pubsubkey)
	return memberhdl.NewHandler(memberServ)
}
