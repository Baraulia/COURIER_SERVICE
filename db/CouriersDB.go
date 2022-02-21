package db

import (
	"database/sql"
	"log"
)

type CourierPostgres struct {
	db *sql.DB
}

func NewCourierPostgres(db *sql.DB) *CourierPostgres {
	return &CourierPostgres{db: db}
}

type Courier struct {
	Id               uint16 `json:"id_courier"`
	CourierName      string `json:"name"`
	ReadyToGo        bool   `json:"ready to go"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Rating           uint16 `json:"rating"`
	Photo            string `json:"photo"`
	Surname          string `json:"surname"`
	NumberOfFailures uint16 `json:"number of failures"`
	Deleted          bool   `json:"deleted"`
}

type SmallInfo struct {
	Id          uint16 `json:"id_courier"`
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	Surname     string `json:"surname"`
}

func (r *CourierPostgres) SaveCourierInDB(courier *Courier) error {
	db := ConnectDB()
	defer db.Close()

	insertValue := `INSERT INTO "couriers" ("name","ready to go","phone_number","email","rating","photo","surname","number of failures","deleted") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := db.Exec(insertValue, courier.CourierName, courier.ReadyToGo, courier.PhoneNumber, courier.Email, courier.Rating, courier.Photo, courier.Surname, courier.NumberOfFailures, courier.Deleted)

	if err != nil {

		log.Println("Error of saving courier in db :" + err.Error())
		return err
	}
	return nil
}

func (r *CourierPostgres) GetCouriersFromDB(Couriers *[]SmallInfo) error {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname" from "couriers"`

	get, err := db.Query(selectValue)

	if err != nil {

		log.Println("Error of getting list of couriers :" + err.Error())
		return err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		*Couriers = append(*Couriers, courier)
	}
	return nil
}

func (r *CourierPostgres) GetCourierFromDB(Couriers *[]SmallInfo, id int) error {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select id_courier,name,phone_number,photo, surname from couriers where id_courier = $1`
	get, err := db.Query(selectValue, id)

	if err != nil {
		log.Println("Error of getting courier :" + err.Error())
		return err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		*Couriers = append(*Couriers, courier)
	}
	return nil
}
