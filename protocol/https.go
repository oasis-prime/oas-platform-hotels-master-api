package protocol

import (
	"encoding/json"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers"
	"github.com/oasis-prime/oas-platform-hotels-master-api/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	gmysql "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

func graphqlHandler() gin.HandlerFunc {
	// oasis-hotel-api
	// Ddg[/Fq&@9J^N4L;
	// var dsn = "oasis-hotel-api@cloudsql(oas-platform:asia-southeast1:oasis-prime)/oasis-master?charset=utf8&parseTime=True&loc=UTC"

	// dsn := "oasis-hotel-api:Ddg[/Fq&@9J^N4L;@tcp(34.124.129.16:3306)/oasis-master?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{

	// })

	// if err != nil {
	// 	panic(err)
	// }
	sqlDB, err := gmysql.Dial("oas-platform:asia-southeast1:oasis-prime", "oasis-trigger-event")
	if err != nil {
		// log.Printf(err.Error())
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		// log.Printf(err.Error())
		panic(err)
	}

	db.AutoMigrate(&repositories.SystemsMember{})

	memberRepo := repositories.NewMemberRepo(db)
	memberServ := services.NewMemberService(memberRepo)
	memberHandler := handlers.NewMemberHandler(memberServ, "")

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				MemberHandler: memberHandler,
			}},
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

	c, _ := json.Marshal(configs.GetConfig())
	fmt.Println(string(c))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":" + configs.GetConfig().App.Port)

	return nil
}
