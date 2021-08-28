package lib

import (
	"fmt"
	"os"
	_ "time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Dialect    = "mysql"
	DBUser     = os.Getenv("MYSQL_USER")
	DBPass     = os.Getenv("MYSQL_PASSWORD")
	DBProtocol = "tcp(db:3306)"
	DBName     = os.Getenv("MYSQL_DATABASE")
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
