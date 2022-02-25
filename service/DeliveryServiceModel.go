package service

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)

type DeliveryService struct {
	repo dao.Repository
}

func NewDeliviveryService(repo dao.Repository) *DeliveryService {
	return &DeliveryService{
		repo: repo,
	}
}
func (s *DeliveryService) CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error) {
	id, err := s.repo.SaveDeliveryServiceInDB(&DeliveryService)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("Error in DeliveryServiceService: %s", err)
	}
	return id, nil
}
