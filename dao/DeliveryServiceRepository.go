package dao

import (
	"database/sql"
	"fmt"
	"log"
)

type DeliveryServicePostgres struct {
	db *sql.DB
}

func NewDeliveryServicePostgres(db *sql.DB) *DeliveryServicePostgres {
	return &DeliveryServicePostgres{db: db}
}

type DeliveryService struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	ManagerId   int    `json:"manager_id"`
	Status      string `json:"status"`
}

func (r *DeliveryServicePostgres) SaveDeliveryServiceInDB(service *DeliveryService) (int, error) {
	row := r.db.QueryRow("INSERT INTO delivery_service (name, email, photo, description, phone_number,manager_id, status) VALUES ($1, $2, $3, $4, $5,$6, $7) RETURNING id", service.Name, service.Email, service.Photo, service.Description, service.PhoneNumber, service.ManagerId, service.Status)
	var id int
	if err := row.Scan(&id); err != nil {
		log.Println(fmt.Sprintf("Create Delivery : error:%s", err))
		return 0, fmt.Errorf("Create Delivery Service: error:%s", err)
	}
	return id, nil
}

func (r *DeliveryServicePostgres) GetDeliveryServiceByIdFromDB(Id int) (DeliveryService, error) {
	var service DeliveryService
	res, err := r.db.Query(fmt.Sprintf("SELECT id, name,email,photo,description,phone_number,manager_id,status FROM delivery_service Where id=%d", Id))
	if err != nil {
		log.Println(err)
		return DeliveryService{}, err
	}
	for res.Next() {
		err = res.Scan(&service.Id, &service.Name, &service.Email, &service.Photo, &service.Description, &service.PhoneNumber, &service.ManagerId, &service.Status)
		if err != nil {
			log.Println(err)
			return DeliveryService{}, err
		}
	}
	return service, nil
}
