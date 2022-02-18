package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func OpenDB() (*sql.DB,error) {
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
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil,err
	}
	return db,err
}


func NewPostgresDB(dbs PostgresDB) (*sql.DB, error) {
	pgsqlConn:=fmt.Sprintf("host= %s port= %s user=%s password=%s dbname=%s sslmode=disable",dbs.Host,dbs.Port, dbs.User, dbs.Password,dbs.DBName)
	db, err :=sql.Open("postgres",pgsqlConn)
	if err != nil {

		return nil, fmt.Errorf("error connecting to database:%s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
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
