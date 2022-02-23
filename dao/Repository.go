package dao

import "database/sql"

type Repository struct {
	DeliveryServiceRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryServicePostgres(db),
	}
}

type DeliveryServiceRep interface {
	SaveDeliveryServiceInDB(service *DeliveryService) (int, error)
}
