package service

import (
	"errors"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)

type OrderService struct {
	repo dao.Repository
}

func NewOrderService(repo dao.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetOrders(id int) ([]dao.Order, error) {
	get, err := s.repo.GetActiveOrdersFromDB(id)
	if get == nil {
		return []dao.Order{}, fmt.Errorf("Error in OrderService: %s", err)
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

func (s *OrderService) GetOrder(id int) (dao.Order, error) {
	get, err := s.repo.GetActiveOrderFromDB(id)
	if (get == dao.Order{}) {
		return dao.Order{}, fmt.Errorf("Error in OrderService: %s", err)
	}
	if err != nil {
		return dao.Order{}, fmt.Errorf("Error with database: %s", err)
	}
	if id == 0 {
		err := errors.New("no id")
		log.Println("id cannot be zero")
		return dao.Order{}, fmt.Errorf("Error in OrderService: %s", err)
	}
	return get, nil
}

func (s *OrderService) ChangeOrderStatus(id uint16) (uint16, error) {
	_, err := s.GetOrder(int(id))
	if err != nil {
		return 0, fmt.Errorf("Error in OrderService: %s", err)
	}
	orderId, err := s.repo.ChangeOrderStatusInDB(id)
	if err != nil {
		return 0, fmt.Errorf("Error with database: %s", err)
	}
	return orderId, nil
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
	fmt.Println(Order)
	return Order, nil
}

func (s *OrderService) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error) {
	var Order = []dao.Order{}
	if limit <= 0 || page <= 0 {
		err := errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil, fmt.Errorf("Error in OrderService: %s", err)
	}
	Order, totalCount := s.repo.GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idService, Month, Year)
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

func (s *OrderService) AssigningOrderToCourier(order dao.Order) error {
	if err := s.repo.AssigningOrderToCourierInDB(order); err != nil {
		log.Println(err)
		return fmt.Errorf("Error in OrderService: %s", err)
	}
	return nil
}

func (s *OrderService) CreateOrder(order *courierProto.OrderCourierServer) error {
	return s.repo.OrderRep.CreateOrder(order)
}

func (s *OrderService) GetServices() (*courierProto.ServicesResponse, error) {
	return s.repo.OrderRep.GetServices()
}