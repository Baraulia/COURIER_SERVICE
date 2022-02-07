package db

import (
	"fmt"
	"strconv"
)

type Courier struct {
	IdCourier        uint16 `json:"id_courier"`
	CourierName      string `json:"courier_name"`
	ReadyToGo        bool   `json:"ready_to_go"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Rating           uint16 `json:"rating"`
	Photo            string `json:"photo"`
	NumberOfFailures uint16 `json:"number_of_failures"`
	Deleted          bool   `json:"deleted"`
}

type SmallInfo struct {
	IdCourier   uint16 `json:"id_courier"`
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
}

func GetCouriersFromDB(Couriers *[]SmallInfo) {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select "id_courier","name", "phone_number","photo" from "couriers"`

	get, err := db.Query(selectValue)

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.IdCourier, &courier.CourierName, &courier.PhoneNumber, &courier.Photo)
		*Couriers = append(*Couriers, courier)
	}
}

func GetOneCourierFromDB(Couriers *SmallInfo, id string) {
	db := ConnectDB()
	defer db.Close()
	l, _ := strconv.Atoi(id)

	get, err := db.Query(fmt.Sprintf("Select id_courier,name,phone_number,photo from couriers where id_courier = %d", l))

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.IdCourier, &courier.CourierName, &courier.PhoneNumber, &courier.Photo)
		*Couriers = courier
	}
}
