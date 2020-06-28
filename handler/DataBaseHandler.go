package handler

import (
	"fmt"

	"github.com/SoundRequest/OAuth2Server/structure"
	"github.com/jinzhu/gorm"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// InitDB DataBase Connection
func InitDB(dbtype string, connectionInfo string) (*gorm.DB, error) {

	db, err := gorm.Open(dbtype, connectionInfo)
	fmt.Println("Successed To Connect Database")

	fmt.Println("Performing AutoMigrate...")
	var models = []interface{}{&structure.Authentication{}}
	db.AutoMigrate(models...)
	fmt.Println("Successfully performed AutoMigrate")

	return db, err
}
