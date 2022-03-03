package service

import (
	"github.com/Baraulia/COURIER_SERVICE/dao"
)

//go:generate mockgen -source=Service.go -destination=mocks/mock.go

type OrderApp interface {
	GetOrder(id int) (dao.Order, error)
	ChangeOrderStatus(id uint16) (uint16, error)
}

type Service struct {
	OrderApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewOrderService(*rep),
	}
}
