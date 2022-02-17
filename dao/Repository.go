package dao

import "database/sql"

type Repository struct {
	OrderRep
	DeliveryServiceRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewOrderPostgres(db),
		NewDeliveryServicePostgres(db),
	}
}

type OrderRep interface {
	GetCourierCompletedOrdersWithPage_fromDB(limit,page,idCourier int) ([]Order,int)
	GetAllOrdersOfCourierServiceWithPage_fromDB(limit,page,idService int) ([]Order,int)
	GetCourierCompletedOrdersByMouthWithPage_fromDB(limit,page,idCourier,Month int) ([]Order,int)
	AssigningOrderToCourier_InDB(order Order) error
}


type DeliveryServiceRep interface {
	SaveDeliveryService_InDB(service *DeliveryService) (*DeliveryService,error)
}