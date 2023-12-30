package database

import (
	"database/sql"
	"log"
)

func MysqlConnection(conn string) (DB *sql.DB) {
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("Error connecting to database")
		panic(err.Error())
	}
	return DB
}