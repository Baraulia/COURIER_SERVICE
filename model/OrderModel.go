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


func (s *OrderService) AssigningOrderToCourier(order dao.Order) error {
	if err := s.repo.AssigningOrderToCourier_InDB(order); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in OrderService: %s", err)
	}
	return nil
}

func (s *OrderService) GetAllServiceCompletedOrders(limit, page, idService int) ([]dao.DetailedOrder, error){
	var Order = []dao.DetailedOrder{}

	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.GetAllServiceCompletedOrders_fromDB(limit, page, idService)
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


func (s *OrderService) GetDetailedOrdersById(idOrder int) (interface{}, error){
	var Order dao.DetailedOrder
	Order, err :=s.repo.GetDetailedOrdersById_FromDB(idOrder)
	if err!=nil{
		err := errors.New("no id")
		log.Println("no more id")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return Order,nil
}