package service

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"log"
)

type CourierService struct {
	repo dao.Repository
}

func NewCourierService(repo dao.Repository) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) GetCouriers() ([]model.SmallInfo, error) {
	get, err := s.repo.CourierRep.GetCouriersFromDB()
	if get == nil {
		return []model.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	return get, nil
}

func (s *CourierService) GetCourier(id uint16) (model.SmallInfo, error) {
	get, err := s.repo.CourierRep.GetCourierFromDB(id)
	if (get == nil) {
		return model.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return model.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	return *get, nil
}

func (s *CourierService) SaveCourier(courier *model.Courier) (uint16, error) {
	return s.repo.CourierRep.SaveCourierInDB(courier)
}
