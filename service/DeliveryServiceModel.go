package service

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)

type DeliveryService struct {
	repo dao.Repository
}

func NewDeliveryService(repo dao.Repository) *DeliveryService {
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

func (s *DeliveryService) GetDeliveryServiceById(Id int) (dao.DeliveryService, error) {
	var service dao.DeliveryService
	service, err := s.repo.GetDeliveryServiceByIdFromDB(Id)
	if err != nil {
		err := errors.New("no id")
		log.Println("no more id")
		return dao.DeliveryService{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return service, nil
}

func (s *DeliveryService) GetAllDeliveryServices() ([]dao.DeliveryService, error) {
	var services = []dao.DeliveryService{}
	services, err := s.repo.GetAllDeliveryServicesFromDB()
	if err != nil {
		log.Println(err)
		return []dao.DeliveryService{}, fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return services, nil
}

func (s *DeliveryService) UpdateDeliveryService(service dao.DeliveryService) error {
	if err := s.repo.UpdateDeliveryServiceInDB(service); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in DeliveryService: %s", err)
	}
	return nil
}
