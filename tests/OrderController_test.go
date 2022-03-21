package tests

import (
	"bytes"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/Baraulia/COURIER_SERVICE/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetOrders(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, courier dao.Order)

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
		name                   string
		inputBody              string
		inputCourier           dao.Order
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehavior           mockBehavior
		mockBehaviorCheck      mockBehaviorCheck
		expectedStatusCode     int
		expectedRequestBody    string
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
			mockBehavior: func(s *mock_service.MockAllProjectApp, courier dao.Order) {
				s.EXPECT().GetOrders(3).Return(orders, nil)
			},
			inputPerms: "",
			inputRole:  "Courier",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier", "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)
			testCase.mockBehaviorParseToken(get, testCase.inputToken)
			testCase.mockBehaviorCheck(get, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: get}
			handler := controller.NewHandler(services)

			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/orders/3", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_GetOneOrder(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, courier dao.Order)

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
		name                   string
		inputBody              string
		inputCourier           dao.Order
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehavior           mockBehavior
		mockBehaviorCheck      mockBehaviorCheck
		expectedStatusCode     int
		expectedRequestBody    string
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
			mockBehavior: func(s *mock_service.MockAllProjectApp, courier dao.Order) {
				s.EXPECT().GetOrder(1).Return(ord, nil)
			},
			inputPerms: "",
			inputRole:  "Courier",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier", "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"11.11.2022","restaurant_address":"","picked":false}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)
			testCase.mockBehaviorParseToken(get, testCase.inputToken)
			testCase.mockBehaviorCheck(get, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: get}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order/1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			//assert.Equal(t, testCase.expectedRequestBody,w.Body.String())
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}

func TestHandler_UpdateOrder(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, order dao.Order)

	testTable := []struct {
		name                   string
		inputBody              string
		inputOrder             dao.Order
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehavior           mockBehavior
		mockBehaviorCheck      mockBehaviorCheck
		expectedStatusCode     int
	}{
		{
			name:      "OK",
			inputBody: `{"courier_id":8}`,
			inputOrder: dao.Order{
				Id:        1,
				IdCourier: 8,
			},
			mockBehavior: func(s *mock_service.MockAllProjectApp, order dao.Order) {
				s.EXPECT().AssigningOrderToCourier(order).Return(nil)
			},
			inputPerms: "",
			inputRole:  "Courier",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier", "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode: 204,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)
			testCase.mockBehaviorParseToken(get, testCase.inputToken)
			testCase.mockBehaviorCheck(get, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: get}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/orders/1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}

func TestHandler_GetDetailedOrdersById(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, order *dao.DetailedOrder)

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
		CourierSurname:     "",
		CourierPhoneNumber: "1234567",
	}

	testTable := []struct {
		name                   string
		inputBody              string
		inputOrder             dao.DetailedOrder
		inputPerms             string
		inputRole              string
		inputToken             string
		mockBehaviorParseToken mockBehaviorParseToken
		mockBehavior           mockBehavior
		mockBehaviorCheck      mockBehaviorCheck
		expectedStatusCode     int
		expectedRequestBody    string
	}{
		{
			name:      "OK",
			inputBody: `{"id":1}`,
			inputOrder: dao.DetailedOrder{
				IdOrder: 1,
			},
			mockBehavior: func(s *mock_service.MockAllProjectApp, order *dao.DetailedOrder) {
				s.EXPECT().GetDetailedOrderById(1).Return(ord, nil)
			},
			inputPerms: "",
			inputRole:  "Courier",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier", "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"delivery_service_id":1,"id":1,"courier_id":1,"delivery_time":"2022-02-19T13:34:53.000093589Z","customer_address":"Some address","status":"ready to delivery","order_date":"2022-11-11","restaurant_address":"Some address","picked":true,"name":"Sam","surname":"","phone_number":"1234567"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(get, &testCase.inputOrder)
			testCase.mockBehaviorParseToken(get, testCase.inputToken)
			testCase.mockBehaviorCheck(get, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: get}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/order/detailed/1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
