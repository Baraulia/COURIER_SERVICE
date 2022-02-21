package Models

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"log"
)

type CourierService struct {
	repo db.Repository
}

func NewCourierService(repo db.Repository) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) GetCouriers() ([]db.SmallInfo, error) {
	var Couriers []db.SmallInfo
	err := s.repo.GetCouriersFromDB(&Couriers)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return Couriers, nil
}

func (s *CourierService) GetCourier(id int) ([]db.SmallInfo, error) {
	var Courier []db.SmallInfo
	s.repo.GetCourierFromDB(&Courier, id)
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return nil, fmt.Errorf("Error in CourierService: %s", err)
	}
	return Courier, nil
}

func (s *CourierService) SaveCourier(courier *db.Courier) (*db.Courier, error) {

	err := s.repo.SaveCourierInDB(courier)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return courier, nil
}
