package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func RunningDataBase() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "13781378",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "blogdb",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr.Error())
		return nil
	}

	log.Print("database connected successfully. 😃")
	return db
}
