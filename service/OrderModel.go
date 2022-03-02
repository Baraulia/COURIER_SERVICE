package service

import (
	"errors"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPCserver"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"log"
)

type OrderService struct {
	repo dao.Repository
}

func NewOrderService(repo dao.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrders(id int) ([]model.Order, error) {
	get, err := s.repo.OrderRep.GetActiveOrdersFromDB(id)
	if get == nil {
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return get, nil
}

func (s *OrderService) GetOrder(id int) (*model.Order, error) {
	get, err := s.repo.OrderRep.GetActiveOrderFromDB(id)
	if get == nil {
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return get, nil
}


func (s *OrderService) ChangeOrderStatus(status string, id uint16) (uint16, error) {
	_, err := s.GetOrder(int(id))
	if err != nil {
		return 0, fmt.Errorf("Error in OrderService: %s", err)
	}
	orderId, err := s.repo.ChangeOrderStatusInDB(status, id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return orderId, nil
}

func (s *OrderService) GetCourierCompletedOrders(limit, page, idCourier int) ([]model.DetailedOrder, error) {
	var Order []model.DetailedOrder

	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.OrderRep.GetCourierCompletedOrdersWithPageFromDB(limit, page, idCourier)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return Order, nil
}

func (s *OrderService) GetAllOrdersOfCourierService(limit, page, idService int) ([]model.Order, error) {
	var Order []model.Order
	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.OrderRep.GetAllOrdersOfCourierServiceWithPageFromDB(limit, page, idService)
	LimitOfPages := (totalCount / limit) + 1
	if LimitOfPages < page {
		err := errors.New("no page")
		log.Println("no more pages")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	return Order, nil
}

func (s *OrderService) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]model.Order, error) {
	var Order []model.Order
	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.OrderRep.GetCourierCompletedOrdersByMouthWithPageFromDB(limit, page, idService, Month, Year)
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

	return Order, nil
}

func (s *OrderService) CreateOrder(order *courierProto.OrderCourierServer) error {
	return s.repo.OrderRep.CreateOrder(order)
}

func (s *OrderService) GetServices() (*courierProto.ServiceResponse, error) {
	return s.repo.OrderRep.GetServices()
}
