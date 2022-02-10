package db

import "database/sql"

type Repository struct {
	CourierRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewCourierPostgres(db),
	}
}

type CourierRep interface {
	GetCouriersFromDB(Couriers *[]SmallInfo) error
	GetOneCourierFromDB(Couriers *[]SmallInfo, id int) error
}
