package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"net/http"
	"strconv"
)



     var Orders []dao.Order
func (h *Handler) GetCourierCompletedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such page")
		return
	}
	limit,er1:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er1!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such limit")
		return

	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such id of courier")
		return
	}
	if er==nil && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrders(limit,page,idCourier)
		if err!=nil{
			RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",err))
			return
		}
	json.NewEncoder(w).Encode(Orders) }else {
		RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",er))
		return
	}
}


func (h *Handler) GetAllOrdersOfCourierService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such page")
		return
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such limit")
		return
	}
	idService,er:= strconv.Atoi(r.URL.Query().Get("iddeliveryservice"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such id of delivery service")
		return
	}
	if er==nil && page!=0 && limit!=0 {
		Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService( limit, page, idService)
		if err != nil {
			RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",err))
			return
		}
		json.NewEncoder(w).Encode(Orders)
	}else {
		RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",er))
		return
	}
}

func (h *Handler) GetCourierCompletedOrdersByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such page")
		return
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such limit")
		return
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		RespondWithError(w,http.StatusBadRequest,"no such id of courier")
		return
	}
	Month,er:= strconv.Atoi(r.URL.Query().Get("month"))
	if er!=nil {
		RespondWithError(w,http.StatusBadRequest,"enter number from 1 to 12")
		return
	}
	fmt.Println(limit,page,idCourier,Month)
	if er==nil && Month!=0 && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(	limit,page,idCourier,Month)
		if err!=nil{
			RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",err))
			return
		}
		json.NewEncoder(w).Encode(Orders) }else {
		RespondWithError(w,http.StatusBadRequest,fmt.Sprintf("Error: %s",er))
		return
	}
}