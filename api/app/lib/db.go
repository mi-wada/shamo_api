package lib

import (
	"fmt"
	"os"
	_ "time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBUsername = os.Getenv("MYSQL_USERNAME")
	DBPassword = os.Getenv("MYSQL_PASSWORD")
	DBHost     = os.Getenv("MYSQL_HOST")
	DBPort     = os.Getenv("MYSQL_PORT")
	DBName     = os.Getenv("MYSQL_DATABASE")
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUsername, DBPassword, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
