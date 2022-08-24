package protocol

import (
	"github.com/oasis-prime/oas-platform-core/http/hotelbedshttp"
	"github.com/oasis-prime/oas-platform-core/repositories/customerrepo"
	"github.com/oasis-prime/oas-platform-core/repositories/hotelrepo"
	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelbedsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/hotelsserv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services/memberserv"
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
		memberServ := memberserv.NewService(memberRepo)
		memberHandler = memberhdl.NewHandler(memberServ, "")
	}

	var hotelsHandler *hotelshdl.Handler
	{
		apprequest := hotelbedshttp.NewRequester(con.Hotelbeds.Key, con.Hotelbeds.Secret, con.Hotelbeds.Format)
		hotelsRepo := hotelrepo.NewHotelsRepo(db)
		hotelbedsContentHttp := hotelbedshttp.NewHotelbedsContentHTTP(con.Hotelbeds.Endpoint, apprequest)
		hotelbedsServ := hotelbedsserv.NewService(hotelbedsContentHttp)
		hotelsServ := hotelsserv.NewService(hotelsRepo)
		hotelsHandler = hotelshdl.NewHandler(hotelbedsServ, hotelsServ)
	}

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					MemberHandler: memberHandler,
					HotelsHandler: hotelsHandler,
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
