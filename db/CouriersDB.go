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
	CourierName      string `json:"courier_name"`
	ReadyToGo        bool   `json:"ready_to_go"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Rating           uint16 `json:"rating"`
	Photo            string `json:"photo"`
	Surname          string `json:"surname"`
	NumberOfFailures uint16 `json:"number_of_failures"`
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

	insertValue := `INSERT INTO "couriers" ("name","ready to go","phone_number","email","photo","surname") VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := r.db.Exec(insertValue, courier.CourierName, courier.ReadyToGo, courier.PhoneNumber, courier.Email, courier.Photo, courier.Surname)

	if err != nil {

		log.Println("Error of saving courier in db :" + err.Error())
		return err
	}
	return nil
}

func (r *CourierPostgres) GetCouriersFromDB() ([]SmallInfo, error) {
	var Couriers []SmallInfo

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname" from "couriers"`

	get, err := r.db.Query(selectValue)

	if err != nil {

		log.Println("Error of getting list of couriers :" + err.Error())
		return []SmallInfo{}, err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		Couriers = append(Couriers, courier)
	}
	return Couriers, nil
}

func (r *CourierPostgres) GetCourierFromDB(id int) (SmallInfo, error) {
	var cour SmallInfo

	selectValue := `Select id_courier,name,phone_number,photo, surname from couriers where id_courier = $1`

	get, err := r.db.Query(selectValue, id)

	if err != nil {
		log.Println("Error of getting courier :" + err.Error())
		return SmallInfo{}, err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		cour = courier
	}
	return cour, nil
}
