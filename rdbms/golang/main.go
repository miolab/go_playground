package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		// User:   os.Getenv("DBUSER"),
		// Passwd: os.Getenv("DBPASS"),
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "mysql_container:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	fmt.Println("Started setting...")

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error!(open err)")
		log.Fatal(err)
	}

	fmt.Println("Opened DB...")

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error!(ping err)")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
