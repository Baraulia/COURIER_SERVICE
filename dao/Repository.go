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
	GetCourierCompletedOrdersWithPage_fromDB(Orders *[]Order,limit,page,idCourier int) int
	GetAllOrdersOfCourierServiceWithPage_fromDB(Orders *[]Order,limit,page,idService int) int
	GetCourierCompletedOrdersByMouthWithPage_fromDB(Orders *[]Order,limit,page,idCourier,Month int) int
}
