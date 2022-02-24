package service

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
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
