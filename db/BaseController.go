package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var dbConnection *gorm.DB

// InitDB DataBase Connection
func InitDB(dbtype string, connectionInfo string) (*gorm.DB, error) {
	var err error
	dbConnection, err = gorm.Open(dbtype, connectionInfo)

	fmt.Println("Successed To Connect Database")
	fmt.Println("Connection To " + connectionInfo)

	return dbConnection, err
}

// DB returns Database Connection
func DB() *gorm.DB {
	return dbConnection
}
