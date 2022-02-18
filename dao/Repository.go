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
	SaveDeliveryService_InDB(service *DeliveryService) (*DeliveryService, error)
}
