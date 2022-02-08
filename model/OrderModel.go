package model

import (
	"CourierService/dao"
	"errors"
	"fmt"
)

func GetCourierCompletedOrders( limit,page,idCourier int) ([]dao.Order,error){
	var Order = []dao.Order{}
	totalCount:=dao.GetCourierCompletedOrdersWithPage_fromDB(&Order,limit,page,idCourier)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		return nil,errors.New("No more pages")
	}
	if limit==0 || page==0 {
		return nil,errors.New("no page or limit")
	}
	fmt.Println(Order)
	return Order,nil
}


func GetAllOrdersOfCourierService(limit,page,idService int) ([]dao.Order,error){
	var Order = []dao.Order{}
	totalCount:=dao.GetAllOrdersOfCourierServiceWithPage_fromDB(&Order,limit,page,idService)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		return nil,errors.New("No more pages")
	}
	if limit==0 || page==0 {
		//dao.Get_books_fronDB(&AllBooks)
		//log.Println("this is all books")
		return nil,errors.New("no page or limit")
	}
	fmt.Println(Order)
	return Order,nil
}

func GetCourierCompletedOrdersByMonth(limit,page,idService,Month int) ([]dao.Order,error){
	var Order = []dao.Order{}
	totalCount:=dao.GetCourierCompletedOrdersByMouthWithPage_fromDB(&Order,limit,page,idService,Month)
	LimitOfPages:=(totalCount/limit)+1
	if LimitOfPages<page{
		return nil,errors.New("No more pages")
	}
	if limit==0 || page==0 {
		return nil,errors.New("no page or limit")
	}
	if Month>=13 || Month<1{
		return nil,errors.New("enter correct month")
	}
	return Order,nil
}
