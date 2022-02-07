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

func GetCouriersFromDB(Couriers *[]Courier) {
	db := ConnectDB()
	defer db.Close()

	selectValue := `Select * from "couriers"`

	get, err := db.Query(selectValue)

	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var courier Courier
		err = get.Scan(&courier.IdCourier, &courier.CourierName, &courier.ReadyToGo, &courier.PhoneNumber, &courier.Email, &courier.Rating, &courier.Photo, &courier.NumberOfFailures, &courier.Deleted)
		*Couriers = append(*Couriers, courier)
	}
}
