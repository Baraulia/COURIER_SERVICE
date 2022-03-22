package service

import (
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	authProto "github.com/Baraulia/COURIER_SERVICE/GRPCC"
	"github.com/Baraulia/COURIER_SERVICE/GRPCC/grpcClient"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type AllProjectApp interface {
	GetOrder(id int) (dao.Order, error)
	GetOrders(id int) ([]dao.Order, error)
	ChangeOrderStatus(text string, id uint16) (uint16, error)
	GetOrderForChange(id int) (dao.Order, error)
	GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error)
	GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.DetailedOrder, error)
	GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error)
	AssigningOrderToCourier(order dao.Order) error
	GetDetailedOrderById(Id int) (*dao.DetailedOrder, error)
	CreateOrder(order *courierProto.OrderCourierServer) (*emptypb.Empty, error)
	GetServices(in *emptypb.Empty) (*courierProto.ServicesResponse, error)
	GetCompletedOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error)
	GetCompletedOrdersOfCourierServiceByDate(limit, page, idService int) ([]dao.Order, error)
	GetCompletedOrdersOfCourierServiceByCourierId(limit, page, idService int) ([]dao.Order, error)
	GetOrdersOfCourierServiceForManager(limit, page, idService int) ([]dao.DetailedOrder, error)

	GetCouriers() ([]dao.SmallInfo, error)
	GetCourier(id int) (dao.SmallInfo, error)
	SaveCourier(courier *dao.Courier) (*dao.Courier, error)
	UpdateCourier(id uint16) (uint16, error)
	SaveCourierPhoto(cover []byte, id int) error
	GetCouriersOfCourierService(limit, page, id int) ([]dao.Courier, error)

	CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error)
	GetDeliveryServiceById(Id int) (*dao.DeliveryService, error)
	GetAllDeliveryServices() ([]dao.DeliveryService, error)
	UpdateDeliveryService(service dao.DeliveryService) error
	SaveLogoFile(cover []byte, id int) error

	ParseToken(token string) (*authProto.UserRole, error)
	CheckRoleRights(neededPerms []string, neededRole1 string, neededRole2 string, givenPerms string, givenRole string) error
}

type Service struct {
	AllProjectApp
}

func NewService(rep *dao.Repository, grpcCli *grpcClient.GRPCClient) *Service {
	return &Service{
		NewProjectService(*rep, grpcCli),
	}
}
