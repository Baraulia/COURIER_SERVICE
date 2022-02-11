package model

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"log"
)
type OrderService struct {
	repo dao.Repository
}


func NewOrderService(repo dao.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetCourierCompletedOrders( limit,page,idCourier int) ([]dao.Order,error){
	var Order = []dao.Order{}
	if limit<=0 || page<=0 {
		err:=errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	Order,totalCount:=s.repo.GetCourierCompletedOrdersWithPage_fromDB(limit,page,idCourier)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page {
		err:=errors.New("no page")
		log.Println("no more pages")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	fmt.Println(Order)
	return Order,nil
}


func (s *OrderService) GetAllOrdersOfCourierService(limit,page,idService int) ([]dao.Order,error){
	var Order = []dao.Order{}
	if limit<=0 || page<=0 {
		err:=errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	totalCount:=s.repo.GetAllOrdersOfCourierServiceWithPage_fromDB(&Order,limit,page,idService)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		err:=errors.New("no page")
		log.Println("no more pages")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	fmt.Println(Order)
	return Order,nil
}

func (s *OrderService) GetCourierCompletedOrdersByMonth(limit,page,idService,Month int) ([]dao.Order,error){
	var Order = []dao.Order{}
	if limit<=0 || page<=0 {
		err:=errors.New("no page or limit")
		log.Println("no more pages or limit")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	totalCount:=s.repo.GetCourierCompletedOrdersByMouthWithPage_fromDB(&Order,limit,page,idService,Month)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		err:=errors.New("no page")
		log.Println("no more pages")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	if Month>=13 || Month<1{
		err:=errors.New("enter correct month")
		log.Println("enter correct month")
		return nil,fmt.Errorf("Error in OrderService: %s",err)
	}
	return Order,nil
}
