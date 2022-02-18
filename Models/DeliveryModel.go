package Models

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"log"
)

type DeliveryService struct {
	repo db.Repository
}

func NewDeliveryService(repo db.Repository) *DeliveryService {
	return &DeliveryService{repo: repo}
}

func (s *DeliveryService) GetOrders() ([]db.Order, error) {
	var Orders []db.Order
	err := s.repo.GetActiveOrdersFromDB(&Orders)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return Orders, nil
}

func (s *DeliveryService) GetOrder(id int) ([]db.Order, error) {
	var Order []db.Order
	err := s.repo.GetActiveOrderFromDB(&Order, id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return nil, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return Order, nil
}
