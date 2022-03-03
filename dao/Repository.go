package dao

import "database/sql"

type Repository struct {
	OrderRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryPostgres(db),
	}
}

type OrderRep interface {
	GetActiveOrderFromDB(id int) (Order, error)
	ChangeOrderStatusInDB(id uint16) (uint16, error)
}
