package test

import (
	"bytes"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/model"
	mocks "github.com/Baraulia/COURIER_SERVICE/model/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)


func TestHandler_GetCourierCompletedOrders(t *testing.T) {
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
			//inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputBody: `{"name":"Test","courier_id":1}`,
			inputOrder: []dao.Order{
				{
				//IdDeliveryService: 1,
				//IdOrder:           1,
				IdCourier:         1,
				//DeliveryTime:      "15:00",
				//CustomerAddress:   "Some address",
				//Status:            "ready to delivery",
				//OrderDate:         "11.11.2022",
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetCourierCompletedOrders(1,1,1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}]`+"\n",
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

			r := mux.NewRouter()

			r.HandleFunc("/orders/completed", handler.GetCourierCompletedOrders).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/completed?limit=1&page=1&idcourier=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
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
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "15:00",
					CustomerAddress:   "Some address",
					Status:            "ready to delivery",
					OrderDate:         "11.11.2022",
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetAllOrdersOfCourierService(1,1,1).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}]`+"\n",
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

			r := mux.NewRouter()

			r.HandleFunc("/orders", handler.GetAllOrdersOfCourierService).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders?limit=1&page=1&iddeliveryservice=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
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
			inputBody: `{"name":"Test","delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}`,
			inputOrder: []dao.Order{
				{
					IdDeliveryService: 1,
					IdOrder:           1,
					IdCourier:         1,
					DeliveryTime:      "15:00",
					CustomerAddress:   "Some address",
					Status:            "ready to delivery",
					OrderDate:         "11.11.2022",
				},
			},
			mockBehavior: func(s *mocks.MockOrderApp, order []dao.Order) {
				s.EXPECT().GetCourierCompletedOrdersByMonth(1,1,1,11).Return(orders, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"15:00","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022"}]`+"\n",
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

			r := mux.NewRouter()

			r.HandleFunc("/orders/bymonth", handler.GetCourierCompletedOrdersByMonth).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/bymonth?limit=1&page=1&idcourier=1&month=11", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

