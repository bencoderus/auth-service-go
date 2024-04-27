package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connection *gorm.DB

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	connection = db

	return db
}

func GetDBConnection() *gorm.DB {
	return connection
}
