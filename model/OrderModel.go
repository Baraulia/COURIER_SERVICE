package model

import (
	"fmt"
	"log"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	_ "stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)

type OrderService struct {
	repo dao.Repository
}

func NewOrderService(repo dao.Repository) *OrderService {
	return &OrderService{repo: repo}
}


func (s *OrderService) AssigningOrderToCourier(order dao.Order) error {
	if err := s.repo.AssigningOrderToCourier_InDB(order); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in OrderService: %s", err)
	}
	return nil
}

