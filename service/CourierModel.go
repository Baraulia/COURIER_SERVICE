package service

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)

type CourierService struct {
	repo dao.Repository
}

func NewCourierService(repo dao.Repository) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) GetCouriers() ([]dao.SmallInfo, error) {
	get, err := s.repo.GetCouriersFromDB()
	if get == nil {
		return []dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	return get, nil
}

func (s *CourierService) GetCourier(id int) (dao.SmallInfo, error) {
	get, err := s.repo.GetCourierFromDB(id)
	if (get == dao.SmallInfo{}) {
		return dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return dao.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	return get, nil
}

func (s *CourierService) SaveCourier(courier *dao.Courier) (*dao.Courier, error) {

	err := s.repo.SaveCourierInDB(courier)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return courier, nil
}

func (s *CourierService) UpdateCourier(id uint16) (uint16, error) {
	courierId, err := s.repo.CourierRep.UpdateCourierInDB(id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return courierId, nil
}
