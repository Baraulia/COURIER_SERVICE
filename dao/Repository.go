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
}


