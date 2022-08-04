package protocol

import (
	"fmt"
	"os"

	"github.com/oasis-prime/oas-platform-core/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() {
	dbUser := con.Database.User
	dbPwd := con.Database.PWD
	instanceConnectionName := con.Database.Instance.Connection.Name
	dbName := con.Database.Name

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")

	var dbURI string
	if isSet {
		socketDir = "/cloudsql"
		dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	} else {
		dbURI = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPwd, instanceConnectionName, dbName)
	}

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&repositories.SystemsMember{},
		&repositories.Hotels{},
		&repositories.HotelName{},
		&repositories.HotelDescription{},
	)
}
