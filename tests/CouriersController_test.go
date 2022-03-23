package tests

import (
	"bytes"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/Baraulia/COURIER_SERVICE/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetCouriers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCourierApp, courier dao.SmallInfo)
	var couriers []dao.SmallInfo
	cour := dao.SmallInfo{
		Id:          1,
		CourierName: "test",
		PhoneNumber: "1038812",
		Photo:       "my fav photo",
		Surname:     "Shorokhov",
		Deleted:     true,
	}
	couriers = append(couriers, cour)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        dao.SmallInfo
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov","deleted":true}`,
			inputCourier: dao.SmallInfo{
				Id:          1,
				CourierName: "test",
				PhoneNumber: "1038812",
				Photo:       "my fav photo",
				Surname:     "Shorokhov",
			},
			mockBehavior: func(s *mock_service.MockCourierApp, courier dao.SmallInfo) {
				s.EXPECT().GetCouriers().Return(couriers, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov","deleted":true}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockCourierApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{CourierApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/couriers", handler.GetCouriers)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/couriers", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetOneCourier(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCourierApp, courier dao.Courier)
	cour := dao.Courier{
		Id:          1,
		CourierName: "test",
		PhoneNumber: "1038812",
		Photo:       "my fav photo",
		Surname:     "Shorokhov",
		Deleted:     true,
	}

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        dao.Courier
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov","deleted":true}`,
			inputCourier: dao.Courier{
				Id:                1,
				CourierName:       "test",
				PhoneNumber:       "1038812",
				Photo:             "my fav photo",
				Surname:           "Shorokhov",
				DeliveryServiceId: 0,
			},
			mockBehavior: func(s *mock_service.MockCourierApp, courier dao.Courier) {
				s.EXPECT().GetCourier(1).Return(cour, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id_courier":1,"courier_name":"test","ready_to_go":false,"phone_number":"1038812","email":"","rating":0,"photo":"my fav photo","surname":"Shorokhov","number_of_failures":0,"deleted":true,"delivery_service_id":0}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockCourierApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{CourierApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/courier/:id", handler.GetCourier)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/courier/1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetCourierCompletedOrders(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, order []dao.DetailedOrder)
	var orders []dao.DetailedOrder
	ord := dao.DetailedOrder{
		IdDeliveryService:     1,
		IdOrder:               1,
		IdCourier:             1,
		DeliveryTime:          time.Date(2020, time.May, 2, 2, 2, 2, 2, time.UTC),
		CustomerAddress:       "Some address",
		Status:                "ready to delivery",
		OrderDate:             "11.11.2022",
		CourierPhoneNumber:    "",
		CourierName:           "",
		CourierSurname:        "",
		OrderIdFromRestaurant: 0,
		Picked:                false,
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
			mockBehavior: func(s *mock_service.MockOrderApp, order []dao.DetailedOrder) {
				s.EXPECT().GetCourierCompletedOrders(1, 1, 1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","picked":false,"name":"","surname":"","phone_number":"","id_from_restaurant":0}]}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &service.Service{OrderApp: get}
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
	type mockBehavior func(s *mock_service.MockOrderApp, order []dao.DetailedOrder)
	var orders []dao.DetailedOrder
	ord := dao.DetailedOrder{
		IdDeliveryService:     1,
		IdOrder:               1,
		IdCourier:             1,
		DeliveryTime:          time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
		CustomerAddress:       "Some address",
		Status:                "ready to delivery",
		OrderDate:             "2022-11-11",
		RestaurantAddress:     "Some address",
		Picked:                true,
		CourierName:           "Sam",
		CourierSurname:        "",
		OrderIdFromRestaurant: 0,
		CourierPhoneNumber:    "1234567",
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
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false,"id_from_restaurant":0}}`,
			inputOrder: []dao.DetailedOrder{
				{
					IdDeliveryService:  1,
					IdOrder:            1,
					IdCourier:          1,
					DeliveryTime:       time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
					CustomerAddress:    "Some address",
					Status:             "ready to delivery",
					OrderDate:          "2022-11-11",
					RestaurantAddress:  "Some address",
					Picked:             true,
					CourierName:        "Sam",
					CourierSurname:     "",
					CourierPhoneNumber: "1234567",
				},
			},
			mockBehavior: func(s *mock_service.MockOrderApp, order []dao.DetailedOrder) {
				s.EXPECT().GetAllOrdersOfCourierService(1, 1, 1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"2022-11-11","restaurant_address":"Some address","picked":true,"name":"Sam","surname":"","phone_number":"1234567","id_from_restaurant":0}]}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &service.Service{OrderApp: get}
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
	type mockBehavior func(s *mock_service.MockOrderApp, order []dao.Order)
	var orders []dao.Order
	ord := dao.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      time.Date(2020, time.May, 2, 2, 2, 2, 2, time.UTC),
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
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					Id:                1,
					IdCourier:         1,
					DeliveryTime:      time.Date(2020, time.May, 2, 2, 2, 2, 2, time.UTC),
					CustomerAddress:   "Some address",
					Status:            "ready to delivery",
					OrderDate:         "11.11.2022",
					RestaurantAddress: "",
					Picked:            false,
				},
			},
			mockBehavior: func(s *mock_service.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetCourierCompletedOrdersByMonth(1, 1, 1, 11, 2022).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}]}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &service.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/orders/bymonth", handler.GetCourierCompletedOrdersByMonth)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/bymonth?limit=1&page=1&idcourier=1&month=11&year=2022", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
