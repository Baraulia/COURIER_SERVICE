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
	AssigningOrderToCourier_InDB(order Order) error
	GetAllServiceCompletedOrders_fromDB(limit, page, idService int) ([]DetailedOrder, int)
	GetDetailedOrdersById_FromDB(Id int) (DetailedOrder, error)
}


