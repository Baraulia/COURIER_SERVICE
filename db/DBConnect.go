package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "159.223.1.135"
	port     = 5434
	user     = "courierteam1"
	password = "qwerty"
	dbname   = "courier_db"
)

func ConnectDB() *sql.DB {
	psqlconn := fmt.Sprintf("host= %s port= %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		fmt.Println(err)
	}
	return db
}
