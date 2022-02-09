package dao

type Repository struct {
	OrderRep
}

type OrderRep interface {
	GetCourierCompletedOrdersWithPage_fromDB(Orders *[]Order,limit,page,idCourier int) int
	GetAllOrdersOfCourierServiceWithPage_fromDB(Orders *[]Order,limit,page,idService int) int
	GetCourierCompletedOrdersByMouthWithPage_fromDB(Orders *[]Order,limit,page,idCourier,Month int) int
}
