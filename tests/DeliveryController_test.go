package tests

import (
	"bytes"
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/service"
	mock_service "github.com/Baraulia/COURIER_SERVICE/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetOrders(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryApp, courier db.Order)
	var orders []db.Order
	ord := db.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      "15:00",
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        db.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputCourier: db.Order{
				IdDeliveryService: 1,
				Id:                1,
				IdCourier:         1,
				DeliveryTime:      "15:00",
				CustomerAddress:   "Some address",
				Status:            "ready to delivery",
				OrderDate:         "11.11.2022",
			},
			mockBehavior: func(s *mock_service.MockDeliveryApp, courier db.Order) {
				s.EXPECT().GetOrders().Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockDeliveryApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{DeliveryApp: get}
			handler := Controllers.NewHandler(services)

			r := mux.NewRouter()

			r.HandleFunc("/orders", handler.GetOrders).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetOneOrder(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryApp, courier db.Order)
	var orders []db.Order
	ord := db.Order{
		IdDeliveryService: 1,
		Id:                1,
		IdCourier:         1,
		DeliveryTime:      "15:00",
		CustomerAddress:   "Some address",
		Status:            "ready to delivery",
		OrderDate:         "11.11.2022",
	}
	orders = append(orders, ord)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        db.Order
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputCourier: db.Order{
				IdDeliveryService: 1,
				Id:                1,
				IdCourier:         1,
				DeliveryTime:      "15:00",
				CustomerAddress:   "Some address",
				Status:            "ready to delivery",
				OrderDate:         "11.11.2022",
			},
			mockBehavior: func(s *mock_service.MockDeliveryApp, courier db.Order) {
				s.EXPECT().GetOneOrder(1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockDeliveryApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{DeliveryApp: get}
			handler := Controllers.NewHandler(services)

			r := mux.NewRouter()

			r.HandleFunc("/order", handler.GetOneOrder).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order?id=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
