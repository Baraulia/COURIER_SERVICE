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
	WorkingNow  bool   `json:"working_now"`
	Description string `json:"description"`
	Deleted     bool   `json:"deleted"`
}

func (r *DeliveryServicePostgres) SaveDeliveryService_InDB(service *DeliveryService) (*DeliveryService, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Println(fmt.Sprintf("Create Delivery Service: can not starts transaction:%s", err))
		return nil, fmt.Errorf("Create Delivery Service: can not starts transaction:%s", err)
	}
	row := transaction.QueryRow("INSERT INTO delivery_service (name, email, working_now, description, deleted) VALUES ($1, $2, $3, $4, $5) RETURNING id", service.Name, service.Email, service.WorkingNow, service.Description, service.Deleted)
	var newService DeliveryService
	if err := row.Scan(&newService.Id); err != nil {
		log.Println(fmt.Sprintf("Create Delivery Service: can not starts transaction:%s", err))
		return nil, fmt.Errorf("Create Delivery Service: can not starts transaction:%s", err)
	}
	return &newService, transaction.Commit()
}
