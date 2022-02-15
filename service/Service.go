package service

import (
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type DeliveryApp interface {
	GetOneOrder(id int) ([]db.Order, error)
	GetOrders() ([]db.Order, error)
	ChangeOrderStatus(id int) error
}

type CourierApp interface {
	GetCouriers() ([]db.SmallInfo, error)
	GetOneCourier(id int) ([]db.SmallInfo, error)
}

type Service struct {
	DeliveryApp
	CourierApp
}

func NewService(rep *db.Repository) *Service {
	return &Service{
		Models.NewDeliveryService(*rep),
		Models.NewCourierService(*rep),
	}
}
