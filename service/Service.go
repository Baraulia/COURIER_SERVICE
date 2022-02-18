package service

import (
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type DeliveryApp interface {
	GetOrder(id int) ([]db.Order, error)
	GetOrders() ([]db.Order, error)
}

type Service struct {
	DeliveryApp
}

func NewService(rep *db.Repository) *Service {
	return &Service{
		Models.NewDeliveryService(*rep),
	}
}
