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

func (s *DeliveryService) GetOrders(id int) ([]db.Order, error) {
	get, err := s.repo.GetActiveOrdersFromDB(id)
	if get == nil {
		return []db.Order{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return nil, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return get, nil
}

func (s *DeliveryService) GetOrder(id int) (db.Order, error) {
	get, err := s.repo.GetActiveOrderFromDB(id)
	if (get == db.Order{}) {
		return db.Order{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	if err != nil {
		return db.Order{}, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return db.Order{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return get, nil
}

func (s *DeliveryService) ChangeOrderStatus(id uint16) (uint16, error) {
	_, err := s.GetOrder(int(id))
	if err != nil {
		return 0, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	orderId, err := s.repo.ChangeOrderStatusInDB(id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return orderId, nil
}
