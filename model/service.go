package model

import "stlab.itechart-group.com/go/food_delivery/courier_service/dao"

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.Detailedorder, error)
	GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error)
	GetCourierCompletedOrdersByMonth(limit, page, idService, Month int) ([]dao.Order, error)
}

type Service struct {
	OrderApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
	}
}
