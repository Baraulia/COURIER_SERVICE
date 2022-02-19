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
	GetActiveOrdersFromDB(Orders *[]Order, id int) error
	GetActiveOrderFromDB(Orders *Order, id int) error
	ChangeOrderStatusInDB(id uint16) (uint16, error)
}

type CourierRep interface {
	GetCouriersFromDB(Couriers *[]SmallInfo) error
	GetCourierFromDB(Couriers *SmallInfo, id int) error
}
