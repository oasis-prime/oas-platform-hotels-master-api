package protocol

import (
	"github.com/oasis-prime/oas-platform-core/http/hotelbedshttp"
	"github.com/oasis-prime/oas-platform-core/repositories"
	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelbedssvc"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/membersvc"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers/hotelbedshdl"
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
		memberRepo := repositories.NewMemberRepo(db)
		memberServ := membersvc.NewService(memberRepo)
		memberHandler = memberhdl.NewHandler(memberServ, "")
	}

	var hotelbedsHandler *hotelbedshdl.Handler
	{
		apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
		hotelsRepo := repositories.NewHotelsRepo(db)
		hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
		hotelbedsServ := hotelbedssvc.NewService(hotelbedsContentHttp, hotelsRepo)
		hotelbedsHandler = hotelbedshdl.NewHandler(hotelbedsServ)
	}

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					MemberHandler:    memberHandler,
					HotelbedsHandler: hotelbedsHandler,
				},
			},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ServeHTTP() error {
	r := gin.Default()
	configs.ConfigsInit()
	con = configs.GetConfig()
	DBInit()

	r.Any("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + configs.GetConfig().App.Port)

	return nil
}
