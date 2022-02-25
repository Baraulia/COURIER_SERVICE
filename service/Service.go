package service

import (
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/model"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetOrder(id int) (*model.Order, error)
	GetOrders(id int) ([]model.Order, error)
	ChangeOrderStatus(status string,id uint16) (uint16, error)
	GetCourierCompletedOrders(limit, page, idCourier int) ([]model.DetailedOrder, error)
	GetAllOrdersOfCourierService(limit, page, idService int) ([]model.Order, error)
	GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]model.Order, error)
}

type CourierApp interface {
	GetCouriers() ([]model.SmallInfo, error)
	GetCourier(id uint16) (*model.SmallInfo, error)
	SaveCourier(courier *model.Courier) (uint16, error)
	DeleteCourier(id uint16) (uint16, error)
}

type Service struct {
	OrderApp
	CourierApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
		NewCourierService(*rep),
	}
}
