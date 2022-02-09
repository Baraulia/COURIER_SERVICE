package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"net/http"
	"strconv"
)



func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

     var Orders []dao.Order
func (h *Handler) GetCourierCompletedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	if er==nil && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrders(limit,page,idCourier)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			JSONError(w,b,http.StatusExpectationFailed)
		}
	json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusExpectationFailed)
	}
}


func (h *Handler) GetAllOrdersOfCourierService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	idService,er:= strconv.Atoi(r.URL.Query().Get("iddeliveryservice"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	if er==nil && page!=0 && limit!=0 {
		Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService( limit, page, idService)
		if err != nil {
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
			JSONError(w,b,http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(Orders)
	}else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusBadRequest)
	}
}

func (h *Handler) GetCourierCompletedOrdersByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	Month,er:= strconv.Atoi(r.URL.Query().Get("month"))
	if er!=nil {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	fmt.Println(limit,page,idCourier,Month)
	if er==nil && Month!=0 && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(	limit,page,idCourier,Month)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			JSONError(w,b,http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusBadRequest)
	}
}