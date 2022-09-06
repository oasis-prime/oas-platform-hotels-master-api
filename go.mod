module github.com/oasis-prime/oas-platform-hotels-master-api

go 1.16

require (
	github.com/99designs/gqlgen v0.17.16
	github.com/vektah/gqlparser/v2 v2.5.0
)

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/google/go-querystring v1.1.0
	github.com/jinzhu/copier v0.3.5
	github.com/joho/godotenv v1.4.0
	github.com/nqd/flat v0.1.1
	github.com/oasis-prime/oas-platform-core v0.1.4
	// github.com/oasis-prime/oas-platform-core v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.12.0
	github.com/valyala/fasthttp v1.39.0
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.8
)

// replace github.com/oasis-prime/oas-platform-core => ../oas-platform-core
