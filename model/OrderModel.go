package model

import (
	"errors"
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

func (s *OrderService) GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error) {
	var Order = []dao.DetailedOrder{}

	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if Order == nil {
		err := errors.New("no id")
		log.Println("no more id")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	fmt.Println(Order)
	return Order, nil
}

func (s *OrderService) GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error) {
	var Order = []dao.Order{}
	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if Order == nil {
		err := errors.New("no id")
		log.Println("no more id")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	fmt.Println(Order)
	return Order, nil
}

func (s *OrderService) GetCourierCompletedOrdersByMonth(limit, page, idService, Month int) ([]dao.Order, error) {
	var Order = []dao.Order{}
	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idService, Month)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if Month >= 13 || Month < 1 {
		err := errors.New("enter correct month")
		log.Println("enter correct month")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if Order == nil {
		err := errors.New("no id")
		log.Println("no more id")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return Order, nil
}
