package model

import (
	"fmt"
	"log"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)

type DelServService struct {
	repo dao.Repository
}

func NewDelivServService(repo dao.Repository) *DelServService{
	return &DelServService{
		repo: repo,
	}
}
func (s *DelServService) CreateDeliveryService(DeliveryService dao.DeliveryService) (*dao.DeliveryService, error){
	var newDeliveryServ *dao.DeliveryService
	newDeliveryServ, err := s.repo.SaveDeliveryService_InDB(&DeliveryService)
	if err != nil {
		log.Println(err)
		return  nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return newDeliveryServ, nil
}
