package protocol

import (
	"github.com/joho/godotenv"
	"github.com/oasis-prime/oas-platform-core/http/hotelbedshttp"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/googleserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelbedsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/memberserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/googlehdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelshdl"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/memberhdl"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

var (
	db  *gorm.DB
	con *configs.Config
)

func graphqlHandler() gin.HandlerFunc {
	var memberHandler *memberhdl.Handler
	{
		memberRepo := customerrepo.NewMemberRepo(db)
		memberServ := memberserv.NewService(memberRepo, con.Google.Projectid, con.Google.Pubsubkey)
		memberHandler = memberhdl.NewHandler(memberServ)
	}

	var hotelbedsHandler *hotelbedshdl.Handler
	{
		apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
		hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
		hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)
		hotelbedsHandler = hotelbedshdl.NewHandler(hotelbedsServ)
	}

	var hotelsHandler *hotelshdl.Handler
	{
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
		hotelsHandler = hotelshdl.NewHandler(hotelsServ)
	}

	var googleHAndler *googlehdl.Handler
	{
		servPlance := googleserv.NewService("AIzaSyDOS6MedhWdvMMyRvGkRTROvTXnf8exZW8")
		googleHAndler = googlehdl.NewHandler(servPlance)
	}

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					MemberHandler:    memberHandler,
					HotelbedsHandler: hotelbedsHandler,
					HotelsHandler:    hotelsHandler,
					GoogleHandler:    googleHAndler,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ServeHTTP() error {
	godotenv.Load()
	r := gin.Default()
	configs.ConfigsInit()
	con = configs.GetConfig()
	DBInit()
	r.Use(CORSMiddleware())
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	r.Run(":" + configs.GetConfig().App.Port)

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
