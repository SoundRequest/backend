package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB DataBase Connection
func InitDB(dbtype string, connectionInfo string) (*sql.DB, error) {
	db, err := sql.Open(dbtype, connectionInfo)

	printVersion(db)
	return db, err
}

func printVersion(db *sql.DB) {
	var version string
	var err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to :", version)
}
