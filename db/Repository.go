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
	GetActiveOrdersFromDB(id int) ([]Order, error)
	GetActiveOrderFromDB(id int) (Order, error)
	ChangeOrderStatusInDB(id uint16) (uint16, error)
}

type CourierRep interface {
	GetCouriersFromDB() ([]SmallInfo, error)
	GetCourierFromDB(id int) (SmallInfo, error)
}
