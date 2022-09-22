package protocol

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBInit() {
	dbUser := con.Database.User
	dbPwd := con.Database.PWD
	instanceConnectionName := con.Database.Instance.Connection.Name
	dbName := con.Database.Name

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPwd, instanceConnectionName, dbName)

	var err error
	db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}
}
