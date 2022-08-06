package protocol

import (
	"fmt"
	"os"

	"github.com/oasis-prime/oas-platform-core/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		dbURI = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, instanceConnectionName, dbName)
	}
	var err error
	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
		AllowGlobalUpdate:                        true,
		Logger:                                   logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&repositories.SystemsMember{},
		&repositories.HotelCity{},
		&repositories.HotelName{},
		&repositories.HotelDescription{},
		&repositories.HotelCoordinates{},
		&repositories.HotelAddress{},
		&repositories.HotelPhone{},
		&repositories.HotelRooms{},
		&repositories.HotelFacility{},
		&repositories.HotelIssues{},
		&repositories.HotelInterestPoints{},
		&repositories.HotelImages{},
		&repositories.Hotels{},
	)
}
