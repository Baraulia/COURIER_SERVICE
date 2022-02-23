package model

import (
	"fmt"
	"log"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)

type DelServService struct {
	repo dao.Repository
}

func NewDelivServService(repo dao.Repository) *DelServService {
	return &DelServService{
		repo: repo,
	}
}
func (s *DelServService) CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error) {
	id, err := s.repo.SaveDeliveryServiceInDB(&DeliveryService)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Error in OrderService: %s", err)
	}
	return id, nil
}
