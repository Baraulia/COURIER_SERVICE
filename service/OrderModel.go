package service

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)

type OrderService struct {
	repo dao.Repository
}

func NewOrderService(repo dao.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrder(id int) (dao.Order, error) {
	get, err := s.repo.GetActiveOrderFromDB(id)
	if (get == dao.Order{}) {
		return dao.Order{}, fmt.Errorf("Error in OrderService: %s", err)
	}
	if err != nil {
		return dao.Order{}, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return dao.Order{}, fmt.Errorf("Error in OrderService: %s", err)
	}
	return get, nil
}

func (s *OrderService) ChangeOrderStatus(id uint16) (uint16, error) {
	_, err := s.GetOrder(int(id))
	if err != nil {
		return 0, fmt.Errorf("Error in OrderService: %s", err)
	}
	orderId, err := s.repo.ChangeOrderStatusInDB(id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return orderId, nil
}
