package dao

import (
	"database/sql"
	"fmt"
	"log"
)

type CourierPostgres struct {
	db *sql.DB
}

func NewCourierPostgres(db *sql.DB) *CourierPostgres {
	return &CourierPostgres{db: db}
}

type Courier struct {
	Id                uint16 `json:"id_courier"`
	CourierName       string `json:"courier_name"`
	ReadyToGo         bool   `json:"ready_to_go"`
	PhoneNumber       string `json:"phone_number"`
	Email             string `json:"email"`
	Rating            uint16 `json:"rating"`
	Photo             string `json:"photo"`
	Surname           string `json:"surname"`
	NumberOfFailures  uint16 `json:"number_of_failures"`
	Deleted           bool   `json:"deleted"`
	DeliveryServiceId uint16 `json:"delivery_service_id"`
}

type SmallInfo struct {
	Id          uint16 `json:"id_courier"`
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	Surname     string `json:"surname"`
	Deleted     bool   `json:"deleted"`
}

func (r *CourierPostgres) SaveCourierInDB(courier *Courier) error {

	insertValue := `INSERT INTO "couriers" ("name","ready to go","phone_number","email","photo","surname", "delivery_service_id") VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := r.db.Exec(insertValue, courier.CourierName, courier.ReadyToGo, courier.PhoneNumber, courier.Email, courier.Photo, courier.Surname, courier.DeliveryServiceId)

	if err != nil {

		log.Println("Error of saving courier in dao :" + err.Error())
		return err
	}
	return nil
}

func (r *CourierPostgres) GetCouriersFromDB() ([]SmallInfo, error) {
	var Couriers []SmallInfo

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname", "deleted" from "couriers"`

	get, err := r.db.Query(selectValue)

	if err != nil {

		log.Println("Error of getting list of couriers :" + err.Error())
		return []SmallInfo{}, err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname, &courier.Deleted)
		Couriers = append(Couriers, courier)
	}
	return Couriers, nil
}

func (r *CourierPostgres) GetCourierFromDB(id int) (SmallInfo, error) {
	var cour SmallInfo

	selectValue := `Select id_courier,name,phone_number,photo, surname, deleted from couriers where id_courier = $1`

	get, err := r.db.Query(selectValue, id)

	if err != nil {
		log.Println("Error of getting courier :" + err.Error())
		return SmallInfo{}, err
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname, &courier.Deleted)
		cour = courier
	}
	return cour, nil
}

func (r *CourierPostgres) UpdateCourierInDB(id uint16) (uint16, error) {

	UpdateValue := `UPDATE couriers SET deleted = $1 WHERE id_courier = $2`
	_, err := r.db.Exec(UpdateValue, "true", id)
	if err != nil {
		log.Println("Error with getting courier by id: " + err.Error())
		return 0, fmt.Errorf("updateCourier: error while scanning:%w", err)
	}
	return id, nil
}

func (r *CourierPostgres) GetCouriersWithServiceFromDB() ([]Courier, error) {
	var Couriers []Courier

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname","delivery_service_id" from "couriers"`

	get, err := r.db.Query(selectValue)

	if err != nil {
		log.Println("Error of getting list of couriers :" + err.Error())
		return []Courier{}, err
	}

	for get.Next() {
		var courier Courier
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname, &courier.DeliveryServiceId)
		Couriers = append(Couriers, courier)
	}
	return Couriers, nil
}

func (r *CourierPostgres) UpdateCourierDB(courier Courier) error {
	var oldCourier Courier
	transaction, err := r.db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(`SELECT id_courier, name, surname, delivery_service_id, email, photo, phone_number
                                  FROM couriers Where id_courier=$1`, courier.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	for res.Next() {
		err = res.Scan(&oldCourier.Id, &oldCourier.CourierName, &oldCourier.Surname, &oldCourier.DeliveryServiceId, &oldCourier.Email,
			&oldCourier.Photo, &oldCourier.PhoneNumber)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if courier.CourierName == "" {
		courier.CourierName = oldCourier.CourierName
	}
	if courier.Email == "" {
		courier.Email = oldCourier.Email
	}
	if courier.Photo == "" {
		courier.Photo = oldCourier.Photo
	}
	if courier.Surname == "" {
		courier.Surname = oldCourier.Surname
	}
	if courier.DeliveryServiceId == 0 {
		courier.DeliveryServiceId = oldCourier.DeliveryServiceId
	}
	if courier.PhoneNumber == "" {
		courier.PhoneNumber = courier.PhoneNumber
	}

	s := `UPDATE couriers SET name=$1, surname=$2, delivery_service_id=$3, email=$4, photo=$5, phone_number=$6 
                            WHERE id_courier = $7`
	log.Println(s)
	insert, err := transaction.Query(s, courier.CourierName, courier.Surname, courier.DeliveryServiceId, courier.Email,
		courier.Photo, courier.PhoneNumber, courier.Id)
	defer insert.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
