package service

import (
	"github.com/Baraulia/COURIER_SERVICE/dao"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetOrder(id int) (dao.Order, error)
	GetOrders(id int) ([]dao.Order, error)
	ChangeOrderStatus(id uint16) (uint16, error)
	GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error)
	GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error)
	GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error)
	AssigningOrderToCourier(order dao.Order) error
	GetDetailedOrdersById(Id int) (dao.DetailedOrder, error)
}

type CourierApp interface {
	GetCouriers() ([]dao.SmallInfo, error)
	GetCourier(id int) (dao.SmallInfo, error)
	SaveCourier(courier *dao.Courier) (*dao.Courier, error)
}

type DeliveryServiceApp interface {
	CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error)
	GetDeliveryServiceById(Id int) (dao.DeliveryService, error)
}

type Service struct {
	OrderApp
	CourierApp
	DeliveryServiceApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
		NewCourierService(*rep), NewDeliveryService(*rep),
	}
}
