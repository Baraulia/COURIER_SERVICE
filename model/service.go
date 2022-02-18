package model

import "stlab.itechart-group.com/go/food_delivery/courier_service/dao"

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	AssigningOrderToCourier(order dao.Order) error
}

type Service struct {
	OrderApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
	}
}
