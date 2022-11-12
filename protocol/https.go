package protocol

import (
	"net/http"

	"cloud.google.com/go/pubsub"
	firebase "firebase.google.com/go/v4"
	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

var (
	db  *gorm.DB
	con *configs.Config
	app *firebase.App
	pub *pubsub.Client
)

func graphqlHandler() gin.HandlerFunc {
	memberHandler := memberHandlerInit()
	hotelbedsHandler := hotelbedsHandlerInit()
	hotelsHandler := hotelsHandlerInit()
	googleHandler := googleHandlerInit()
	paymentHandler := paymentHandlerInit()
	popularHandler := popularHandlerInit()

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					MemberHandler:    memberHandler,
					HotelbedsHandler: hotelbedsHandler,
					HotelsHandler:    hotelsHandler,
					GoogleHandler:    googleHandler,
					PaymentHandler:   paymentHandler,
					PopularHandler:   popularHandler,
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
	r := gin.Default()
	configs.ConfigsInit()
	con = configs.GetConfig()

	DBInit()
	PubSubInit()
	FirebaseInit()

	r.Use(CORSMiddleware())
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	r.GET("/", HealthCheck())
	r.Run(":" + con.App.Port)

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

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlDB, _ := db.DB()
		err := sqlDB.Ping()

		if err != nil {
			logrus.Errorln(err)
			c.JSON(http.StatusBadGateway, gin.H{"status": "error"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	}
}
