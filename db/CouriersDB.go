package db

import (
	"fmt"
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
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
}

func GetCouriersFromDB(Couriers *[]SmallInfo) {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select "name", "phone_number","photo" from "couriers"`

	get, err := db.Query(selectValue)

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier SmallInfo
		err = get.Scan(&courier.CourierName, &courier.PhoneNumber, &courier.Photo)
		*Couriers = append(*Couriers, courier)
	}
}
