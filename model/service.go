package model

import "stlab.itechart-group.com/go/food_delivery/courier_service/dao"

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type DeliveryServiceApp interface {
	CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error)
}

type Service struct {
	DeliveryServiceApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewDelivServService(*rep),
	}
}
