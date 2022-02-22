package dao

import "database/sql"

type Repository struct {
	OrderRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewOrderPostgres(db),
	}
}

type OrderRep interface {
	GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier int) ([]DetailedOrder, int)
	GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService int) ([]Order, int)
	GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month, Year int) ([]Order, int)
}
