package dao

import (
	"database/sql"
	"github.com/Baraulia/COURIER_SERVICE/model"
)
//go:generate mockgen -source=Repository.go -destination=mocks/mock.go

type Repository struct {
	OrderRep
	CourierRep
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewDeliveryPostgres(db),
		NewCourierPostgres(db),
	}
}

type OrderRep interface {
	GetActiveOrdersFromDB(id int) ([]model.Order, error)
	GetActiveOrderFromDB(id int) (*model.Order, error)
	ChangeOrderStatusInDB(status string, id uint16) (uint16, error)
	GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier int) ([]model.DetailedOrder, int)
	GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService int) ([]model.Order, int)
	GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month, Year int) ([]model.Order, int)
}

type CourierRep interface {
	SaveCourierInDB(Courier *model.Courier) (uint16,error)
	GetCouriersFromDB() ([]model.SmallInfo, error)
	GetCourierFromDB(id uint16) (*model.SmallInfo, error)
}
