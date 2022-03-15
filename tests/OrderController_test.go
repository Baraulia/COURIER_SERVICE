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

func TestHandler_GetOrders(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, courier dao.Order)
	var orders []dao.Order
	ord := dao.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        dao.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
			inputCourier: dao.Order{
				IdDeliveryService: 1,
				Id:                1,
				IdCourier:         1,
				DeliveryTime:      time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
				CustomerAddress:   "Some address",
				Status:            "ready to delivery",
				OrderDate:         "11.11.2022",
			},
			mockBehavior: func(s *mock_service.MockOrderApp, courier dao.Order) {
				s.EXPECT().GetOrders(3).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/orders/:id", handler.GetOrders)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/3", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetOneOrder(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, courier dao.Order)

	ord := dao.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
	}

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        dao.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
			inputCourier: dao.Order{
				IdDeliveryService: 1,
				Id:                1,
				IdCourier:         1,
				DeliveryTime:      time.Date(2022, 02, 19, 13, 34, 53, 93589, time.UTC),
				CustomerAddress:   "Some address",
				Status:            "ready to delivery",
				OrderDate:         "11.11.2022",
			},
			mockBehavior: func(s *mock_service.MockOrderApp, courier dao.Order) {
				s.EXPECT().GetOrder(1).Return(ord, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/order/:id", handler.GetOrder)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order/1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_UpdateOrder(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, order dao.Order)
	testTable := []struct {
		name               string
		inputBody          string
		inputOrder         dao.Order
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"courier_id":8}`,
			inputOrder: dao.Order{
				Id:        1,
				IdCourier: 8,
			},
			mockBehavior: func(s *mock_service.MockOrderApp, order dao.Order) {
				s.EXPECT().AssigningOrderToCourier(order).Return(nil)
			},
			expectedStatusCode: 204,
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

			r.PUT("/orders/:id", handler.UpdateOrder)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/orders/1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}

func TestHandler_GetDetailedOrdersById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, order *dao.DetailedOrder)

	ord := &dao.DetailedOrder{
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
		CourierPhoneNumber: "1234567",
	}

	testTable := []struct {
		name                string
		inputBody           string
		inputOrder          dao.DetailedOrder
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id":1}`,
			inputOrder: dao.DetailedOrder{
				IdOrder: 1,
			},
			mockBehavior: func(s *mock_service.MockOrderApp, order *dao.DetailedOrder) {
				s.EXPECT().GetDetailedOrderById(1).Return(ord, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"2022-11-11","restaurant_address":"Some address","picked":true,"name":"Sam","phone_number":"1234567"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockOrderApp(c)
			testCase.mockBehavior(get, &testCase.inputOrder)

			services := &service.Service{OrderApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()

			r.GET("/order/detailed/:id", handler.GetDetailedOrderById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order/detailed/1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetAllCompletedOrdersOfCourierService(t *testing.T) {
	type mockBehavior func(s *mock_service.MockOrderApp, order []dao.Order)
	var orders []dao.Order
	ord := dao.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      time.Date(2020, time.May, 2, 2, 2, 2, 2, time.UTC),
		CustomerAddress:   "Some address",
		Status:            "completed",
		OrderDate:         "2022-02-02",
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
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"completed","order_date":"2022-02-02","restaurant_address":"","picked":false}}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					Id:                1,
					IdCourier:         1,
					DeliveryTime:      time.Date(2020, time.May, 2, 2, 2, 2, 2, time.UTC),
					CustomerAddress:   "Some address",
					Status:            "completed",
					OrderDate:         "2022-02-02",
					RestaurantAddress: "",
					Picked:            false,
				},
			},
			mockBehavior: func(s *mock_service.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetAllOrdersOfCourierService(1, 1, 1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2020-05-02T02:02:02.000000002Z","customer_address":"Some address","status":"completed","order_date":"2022-02-02","restaurant_address":"","picked":false}]}`,
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

			r.GET("/orders/service/completed", handler.GetAllOrdersOfCourierService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/service/completed?limit=1&page=1&iddeliveryservice=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
