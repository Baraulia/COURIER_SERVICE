package dao

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/sirupsen/logrus"
)

type CourierPostgres struct {
	db *sql.DB
}

func NewCourierPostgres(db *sql.DB) *CourierPostgres {
	return &CourierPostgres{db: db}
}

func (r *CourierPostgres) SaveCourierInDB(Courier *model.Courier) (uint16, error) {
	var courierId uint16
	row := r.db.QueryRow("INSERT INTO couriers (name, ready_to_go, phone_number, email, photo , surname) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", Courier.CourierName, Courier.ReadyToGo, Courier.PhoneNumber, Courier.Email, Courier.Photo, Courier.Surname)
	if err := row.Scan(&courierId); err != nil {
		logrus.Errorf("CreateCourier: error while scanning for courierId:%s", err)
		return 0, fmt.Errorf("CreateCourier: error while scanning for courierId:%w", err)
	}
	return courierId, nil
}

func (r *CourierPostgres) GetCouriersFromDB() ([]model.SmallInfo, error) {
	var Couriers []model.SmallInfo

	selectValue := `Select "id_courier","name", "phone_number","photo", "surname" from "couriers"`

	get, err := r.db.Query(selectValue)
	if err != nil {
		logrus.Fatalln("Error of getting list of couriers :", err)
		return nil, fmt.Errorf("Error of getting list of couriers :%w", err)
	}

	for get.Next() {
		var courier model.SmallInfo
		err = get.Scan(&courier.Id, &courier.CourierName, &courier.PhoneNumber, &courier.Photo, &courier.Surname)
		Couriers = append(Couriers, courier)
	}
	return Couriers, nil
}

func (r *CourierPostgres) GetCourierFromDB(id uint16) (*model.SmallInfo, error) {
	var cour model.SmallInfo
	query := "Select id_courier,name,phone_number,photo, surname from couriers where id_courier = $1"
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&cour.Id, &cour.CourierName, &cour.PhoneNumber, &cour.Photo, &cour.Surname); err != nil {
		logrus.Errorf("GetCourierById: error while scanning for role:%s", err)
		return nil, fmt.Errorf("GetCourierById: repository error:%w", err)
	}
	return &cour, nil
}

func (r *CourierPostgres) DeleteCourier(id uint16) (uint16, error) {
	var courierId uint16
	row := r.db.QueryRow("DELETE FROM couriers WHERE id_courier=$1 RETURNING id", id)
	if err := row.Scan(&courierId); err != nil {
		logrus.Errorf("DeleteCourierByID: error while scanning for courierId:%s", err)
		return 0, fmt.Errorf("DeleteCourierByID: error while scanning for courierId:%w", err)
	}
	return courierId, nil
}
