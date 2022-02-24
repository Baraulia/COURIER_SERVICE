package model

import "time"

type Courier struct {
	Id               uint16 `json:"id_courier"`
	CourierName      string `json:"courier_name"`
	ReadyToGo        bool   `json:"ready_to_go"`
	PhoneNumber      string `json:"phone_number"`
	Email            string `json:"email"`
	Rating           uint16 `json:"rating"`
	Photo            string `json:"photo"`
	Surname          string `json:"surname"`
	NumberOfFailures uint16 `json:"number_of_failures"`
	Deleted          bool   `json:"deleted"`
}

type SmallInfo struct {
	Id          uint16 `json:"id_courier"`
	CourierName string `json:"courier_name"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	Surname     string `json:"surname"`
}

type Order struct {
	IdDeliveryService int       `json:"delivery_service_id,omitempty"`
	Id                int       `json:"id"`
	IdCourier         int       `json:"courier_id,omitempty"`
	DeliveryTime      time.Time `json:"delivery_time,omitempty"`
	CustomerAddress   string    `json:"customer_address,omitempty"`
	Status            string    `json:"status"`
	OrderDate         string    `json:"order_date"`
	RestaurantAddress string    `json:"restaurant_address"`
	Picked            bool      `json:"picked"`
}

type DetailedOrder struct {
	IdDeliveryService  int       `json:"delivery_service_id,omitempty"`
	IdOrder            int       `json:"id"`
	IdCourier          int       `json:"courier_id,omitempty"`
	DeliveryTime       time.Time `json:"delivery_time,omitempty"`
	CustomerAddress    string    `json:"customer_address,omitempty"`
	Status             string    `json:"status"`
	OrderDate          string    `json:"order_date,omitempty"`
	RestaurantAddress  string    `json:"restaurant_address,omitempty"`
	Picked             bool      `json:"picked"`
	CourierName        string    `json:"name"`
	CourierPhoneNumber string    `json:"phone_number"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
