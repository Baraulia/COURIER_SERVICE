package dao

import (
	"database/sql"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Repository struct {
	OrderRep
	CourierRep
	DeliveryServiceRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryPostgres(db),
		NewCourierPostgres(db),
		NewDeliveryServicePostgres(db),
	}
}

type OrderRep interface {
	GetActiveOrdersFromDB(id int) ([]Order, error)
	GetActiveOrderFromDB(id int) (Order, error)
	ChangeOrderStatusInDB(id uint16) (uint16, error)
	GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier int) ([]DetailedOrder, int)
	GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService int) ([]Order, int)
	GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month, Year int) ([]Order, int)
	AssigningOrderToCourierInDB(order Order) error
	GetDetailedOrderByIdFromDB(Id int) (*DetailedOrder, error)
	CreateOrder(order *courierProto.OrderCourierServer) (*emptypb.Empty, error)
	GetServices(in *emptypb.Empty) (*courierProto.ServicesResponse, error)
	GetAllCompletedOrdersOfCourierServiceFromDB(limit, page, idService int) ([]Order, int)
	GetAllCompletedOrdersOfCourierServiceByDateFromDB(limit, page, idService int) ([]Order, int)
	GetAllCompletedOrdersOfCourierServiceByCourierIdFromDB(limit, page, idService int) ([]Order, int)
}

type CourierRep interface {
	SaveCourierInDB(Courier *Courier) error
	GetCouriersFromDB() ([]SmallInfo, error)
	GetCourierFromDB(id int) (SmallInfo, error)
	UpdateCourierInDB(id uint16) (uint16, error)
}

type DeliveryServiceRep interface {
	SaveDeliveryServiceInDB(service *DeliveryService) (int, error)
	GetDeliveryServiceByIdFromDB(Id int) (*DeliveryService, error)
	GetAllDeliveryServicesFromDB() ([]DeliveryService, error)
	UpdateDeliveryServiceInDB(service DeliveryService) error
}
