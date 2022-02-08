package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func OpenDB() *sql.DB {
	host     := "159.223.1.135"
	port     := 5434
	user     := "courierteam1"
	password := "qwerty"
	dbname   := "courier_db"
	/*host:="localhost"
	port:=5433
	user:="postgres"
	password:="postgres"
	dbname:="courier_db"*/
	pgsqlConn:=fmt.Sprintf("host= %s port= %d user=%s password=%s dbname=%s sslmode=disable",host,port,user, password,dbname)
	db, err :=sql.Open("postgres",pgsqlConn)
	if err!=nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		return nil
	}
	return db
}