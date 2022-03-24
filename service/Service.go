package service

import (
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetOrder(id int) (dao.Order, error)
	GetOrders(id int) ([]dao.Order, error)
	ChangeOrderStatus(text string, id uint16) (uint16, error)
	GetOrderForChange(id int) (dao.Order, error)
	GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error)
	GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.DetailedOrder, error)
	GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error)
	AssigningOrderToCourier(order dao.Order) error
	GetDetailedOrderById(Id int) (*dao.AllInfoAboutOrder, error)
	CreateOrder(order *courierProto.OrderCourierServer) (*emptypb.Empty, error)
	GetServices(in *emptypb.Empty) (*courierProto.ServicesResponse, error)
	GetCompletedOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error)
	GetCompletedOrdersOfCourierServiceByDate(limit, page, idService int) ([]dao.Order, error)
	GetCompletedOrdersOfCourierServiceByCourierId(limit, page, idService int) ([]dao.Order, error)
	GetOrdersOfCourierServiceForManager(limit, page, idService int) ([]dao.DetailedOrder, error)
}

type CourierApp interface {
	GetCouriers() ([]dao.SmallInfo, error)
	GetCourier(id int) (dao.Courier, error)
	SaveCourier(courier *dao.Courier) (*dao.Courier, error)
	UpdateCourier(id uint16) (uint16, error)
	SaveCourierPhoto(cover []byte, id int) error
	GetCouriersOfCourierService(limit, page, idService int) ([]dao.Courier, error)
}

type DeliveryServiceApp interface {
	CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error)
	GetDeliveryServiceById(Id int) (*dao.DeliveryService, error)
	GetAllDeliveryServices() ([]dao.DeliveryService, error)
	UpdateDeliveryService(service dao.DeliveryService) error
	SaveLogoFile(cover []byte, id int) error
}

type Service struct {
	OrderApp
	CourierApp
	DeliveryServiceApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
		NewCourierService(*rep),
		NewDeliveryService(*rep),
	}
}
