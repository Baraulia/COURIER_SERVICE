package db

import "database/sql"

type Repository struct {
	DeliveryRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryPostgres(db),
	}
}

type DeliveryRep interface {
	GetActiveOrdersFromDB(Orders *[]Order) error
	GetActiveOrderFromDB(Orders *[]Order, id int) error
}
