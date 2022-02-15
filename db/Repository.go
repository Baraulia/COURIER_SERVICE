package db

import "database/sql"

type Repository struct {
	DeliveryRep
	CourierRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryPostgres(db),
		NewCourierPostgres(db),
	}
}

type DeliveryRep interface {
	GetActiveOrdersFromDB(Orders *[]Order) error
	GetActiveOrderFromDB(Orders *[]Order, id int) error
	ChangeOrderStatusInDB(id int) error
}

type CourierRep interface {
	GetCouriersFromDB(Couriers *[]SmallInfo) error
	GetOneCourierFromDB(Couriers *[]SmallInfo, id int) error
}
