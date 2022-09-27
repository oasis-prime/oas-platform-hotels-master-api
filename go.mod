module github.com/oasis-prime/oas-platform-hotels-master-api

go 1.16

require (
	github.com/99designs/gqlgen v0.17.20
	github.com/vektah/gqlparser/v2 v2.5.1
)

require (
	cloud.google.com/go/kms v1.4.0 // indirect
	cloud.google.com/go/pubsub v1.3.1
	firebase.google.com/go/v4 v4.8.0
	github.com/gin-gonic/gin v1.8.1
	github.com/google/go-querystring v1.1.0
	github.com/jinzhu/copier v0.3.5
	github.com/joho/godotenv v1.4.0
	github.com/nqd/flat v0.1.1
	github.com/oasis-prime/oas-platform-core v0.1.9
	github.com/oasis-prime/oas-platform-firebase-core v1.0.0
	// github.com/oasis-prime/oas-platform-core v0.0.0-00010101000000-000000000000
	// github.com/oasis-prime/oas-platform-firebase-core v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.13.0
	github.com/valyala/fasthttp v1.40.0
	google.golang.org/api v0.93.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.9
)

// replace github.com/oasis-prime/oas-platform-core => ../oas-platform-core

// replace github.com/oasis-prime/oas-platform-firebase-core => ../oas-platform-firebase-core
