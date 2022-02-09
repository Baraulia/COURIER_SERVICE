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

/*
db, err := database.NewPostgresDB(database.PostgresDB{
        Host:     os.Getenv("HOST"),
        Port:     os.Getenv("DB_PORT"),
        Username: os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_DATABASE"),
        SSLMode:  os.Getenv("DB_SSL_MODE"),
    })
    if err != nil {
        logger.Panicf("failed to initialize db:%s", err.Error())
    }
и запуск сервера
port := os.Getenv("API_SERVER_PORT")
    logger.Fatal(router.Run(":" + port))
*/