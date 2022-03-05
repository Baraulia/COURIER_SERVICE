package dao

import (
	"database/sql"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
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
	CreateOrder(order *courierProto.OrderCourierServer) error
	GetServices() (*courierProto.ServicesResponse, error)
}

type CourierRep interface {
	SaveCourierInDB(Courier *Courier) error
	GetCouriersFromDB() ([]SmallInfo, error)
	GetCourierFromDB(id int) (SmallInfo, error)
}

type DeliveryServiceRep interface {
	SaveDeliveryServiceInDB(service *DeliveryService) (int, error)
}