package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

//database global
var DB *gorm.DB

func SetupDB() *gorm.DB {

	//db config vars
	var dbHost string = os.Getenv("DB_HOST")
	var dbName string = os.Getenv("DB_NAME")
	var dbUser string = os.Getenv("DB_USERNAME")
	var dbPassword string = os.Getenv("DB_PASSWORD")
	var dbPort string = os.Getenv("DB_PORT")

	//connect to db
	db, dbError := gorm.Open("mysql", dbUser+":"+ dbPassword +"@tcp(" + dbHost+ ":"+ dbPort +")/"+ dbName + "?charset=utf8&parseTime=True&loc=Local")
	if dbError != nil {
		panic("Failed to connect to database")
	}

	//fix for connection timeout
	//see: https://github.com/go-sql-driver/mysql/issues/257
	db.DB().SetMaxIdleConns(0)

	return db
}