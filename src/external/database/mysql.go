package database

import (
	"fmt"

	Models "github.com/otaviobernardes/api-go-gin/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var dbName = "api-go-gin"
var dbPassword = "test@123"

func Init() *gorm.DB {
	return connect()
}

func connect() *gorm.DB {
	var err error
	dsn := "root:" + dbPassword + "@tcp(localhost:3306)/" + dbName + "?parseTime=true&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}

	Db.AutoMigrate(&Models.Student{})
	return Db
}
