package protocol

import (
	"encoding/json"
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/configs"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph"
	"github.com/oasis-prime/oas-platform-hotels-master-api/graph/generated"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/core/services"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/handlers"
	"github.com/oasis-prime/oas-platform-hotels-master-api/internal/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

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
