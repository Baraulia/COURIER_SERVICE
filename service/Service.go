package service

import (
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"google.golang.org/protobuf/types/known/emptypb"
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
	GetDetailedOrderById(Id int) (*dao.DetailedOrder, error)
	CreateOrder(order *courierProto.OrderCourierServer) (*emptypb.Empty, error)
	GetServices(in *emptypb.Empty) (*courierProto.ServicesResponse, error)
}

type CourierApp interface {
	GetCouriers() ([]dao.SmallInfo, error)
	GetCourier(id int) (dao.SmallInfo, error)
	SaveCourier(courier *dao.Courier) (*dao.Courier, error)
	UpdateCourier(id uint16) (uint16, error)
	CheckRights(token string, requiredRole string) (bool, error)
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

func NewService(rep *dao.Repository, grpcCli *grpcClient.GRPCClient) *Service {
	return &Service{
		NewOrderService(*rep, grpcCli),
		NewCourierService(*rep, grpcCli),
		NewDeliveryService(*rep, grpcCli),
	}
}
