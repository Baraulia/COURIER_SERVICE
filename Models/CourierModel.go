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

	get, err := s.repo.GetCouriersFromDB()
	if get == nil {
		return []db.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	return get, nil
}

func (s *CourierService) GetCourier(id int) (db.SmallInfo, error) {
	get, err := s.repo.GetCourierFromDB(id)
	if (get == db.SmallInfo{}) {
		return db.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return db.SmallInfo{}, fmt.Errorf("Error in CourierService: %s", err)
	}
	return get, nil
}
