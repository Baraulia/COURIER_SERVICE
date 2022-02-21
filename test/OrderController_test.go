package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"stlab.itechart-group.com/go/food_delivery/courier_service/controller"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"stlab.itechart-group.com/go/food_delivery/courier_service/model"
	mocks "stlab.itechart-group.com/go/food_delivery/courier_service/model/mocks"
	"testing"
)

func TestHandler_GetCourierCompletedOrders(t *testing.T) {
	type mockBehavior func(s *mocks.MockOrderApp, order []dao.DetailedOrder)
	var orders []dao.DetailedOrder
	ord := dao.DetailedOrder{
		IdDeliveryService:  1,
		IdOrder:            1,
		IdCourier:          1,
		DeliveryTime:       "15:00",
		CustomerAddress:    "Some address",
		Status:             "ready to delivery",
		OrderDate:          "11.11.2022",
		CourierPhoneNumber: "",
		CourierName:        "",
		Picked:             false,
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputOrder          []dao.DetailedOrder
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			//inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputBody: `{"name":"Test","courier_id":1}`,
			inputOrder: []dao.DetailedOrder{
				{
					IdCourier: 1,
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.DetailedOrder) {
				s.EXPECT().GetCourierCompletedOrders(1, 1, 1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","picked":false,"name":"","phone_number":""}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mocks.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &model.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/orders/completed", handler.GetCourierCompletedOrders)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/completed?limit=1&page=1&idcourier=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetAllOrdersOfCourierService(t *testing.T) {
	type mockBehavior func(s *mocks.MockOrderApp, order []dao.Order)
	var orders []dao.Order
	ord := dao.Order{
		IdDeliveryService: 1,
		IdOrder:           1,
		IdCourier:         1,
		DeliveryTime:      "15:00",
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
		RestaurantAddress: "",
		Picked:            false,
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputOrder          []dao.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "15:00",
					CustomerAddress:   "Some address",
					Status:            "ready to delivery",
					OrderDate:         "11.11.2022",
					RestaurantAddress: "",
					Picked:            false,
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetAllOrdersOfCourierService(1, 1, 1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mocks.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &model.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/orders", handler.GetAllOrdersOfCourierService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders?limit=1&page=1&iddeliveryservice=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetCourierCompletedOrdersByMonth(t *testing.T) {
	type mockBehavior func(s *mocks.MockOrderApp, order []dao.Order)
	var orders []dao.Order
	ord := dao.Order{
		IdDeliveryService: 1,
		IdOrder:           1,
		IdCourier:         1,
		DeliveryTime:      "15:00",
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
		RestaurantAddress: "",
		Picked:            false,
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputOrder          []dao.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "15:00",
					CustomerAddress:   "Some address",
					Status:            "ready to delivery",
					OrderDate:         "11.11.2022",
					RestaurantAddress: "",
					Picked:            false,
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetCourierCompletedOrdersByMonth(1, 1, 1, 11).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}]`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mocks.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &model.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/orders/bymonth", handler.GetCourierCompletedOrdersByMonth)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/bymonth?limit=1&page=1&idcourier=1&month=11", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
