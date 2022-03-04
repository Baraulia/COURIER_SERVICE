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
)

func TestHandler_CreateDeliveryService(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryServiceApp, service dao.DeliveryService)
	testTable := []struct {
		name                string
		inputBody           string
		inputService        dao.DeliveryService
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test", "email":"test", "photo":"test", "description": "test","phone_number":"1234567", "status": "active"}`,
			inputService: dao.DeliveryService{
				Name:        "test",
				Email:       "test",
				Photo:       "test",
				Description: "test",
				Status:      "active",
				PhoneNumber: "1234567",
			},
			mockBehavior: func(s *mock_service.MockDeliveryServiceApp, service dao.DeliveryService) {
				s.EXPECT().CreateDeliveryService(service).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Invalid request",
			inputBody:           `{"email": "email",}`,
			inputService:        dao.DeliveryService{},
			mockBehavior:        func(r *mock_service.MockDeliveryServiceApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
		{
			name:                "empty fields",
			inputBody:           `{"email": "email"}`,
			inputService:        dao.DeliveryService{},
			mockBehavior:        func(r *mock_service.MockDeliveryServiceApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"empty fields"}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockDeliveryServiceApp(c)
			tt.mockBehavior(newMock, tt.inputService)
			services := &service.Service{DeliveryServiceApp: newMock}
			handler := controller.NewHandler(services)
			r := gin.New()
			r.POST("/deliveryservice", handler.CreateDeliveryService)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/deliveryservice", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_GetAllDeliveryServices(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryServiceApp, service dao.DeliveryService)
	var servicess []dao.DeliveryService
	serv := dao.DeliveryService{
		Id:          1,
		Name:        "name",
		Email:       "email",
		Photo:       "photo",
		Description: "description",
		PhoneNumber: "123",
		ManagerId:   1,
		Status:      "active",
	}
	servicess = append(servicess, serv)

	testTable := []struct {
		name                string
		inputBody           string
		inputService        dao.DeliveryService
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockDeliveryServiceApp, service dao.DeliveryService) {
				s.EXPECT().GetAllDeliveryServices().Return(servicess, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"id":1,"name":"name","email":"email","photo":"photo","description":"description","phone_number":"123","manager_id":1,"status":"active"}]}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockDeliveryServiceApp(c)
			tt.mockBehavior(newMock, tt.inputService)

			services := &service.Service{DeliveryServiceApp: newMock}
			handler := controller.NewHandler(services)

			r := gin.New()
			r.GET("/deliveryservice/", handler.GetAllDeliveryServices)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/deliveryservice/", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())

		})
	}
}

func TestHandler_GetDeliveryServiceById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryServiceApp, service *dao.DeliveryService)

	serv := &dao.DeliveryService{
		Id:          1,
		Name:        "name",
		Email:       "email",
		Photo:       "photo",
		Description: "description",
		PhoneNumber: "123",
		ManagerId:   1,
		Status:      "active",
	}

	testTable := []struct {
		name                string
		inputBody           string
		inputService        dao.DeliveryService
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id":1}`,
			inputService: dao.DeliveryService{
				Id: 1,
			},
			mockBehavior: func(s *mock_service.MockDeliveryServiceApp, service *dao.DeliveryService) {
				s.EXPECT().GetDeliveryServiceById(1).Return(serv, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"name":"name","email":"email","photo":"photo","description":"description","phone_number":"123","manager_id":1,"status":"active"}`,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockDeliveryServiceApp(c)
			tt.mockBehavior(newMock, &tt.inputService)

			services := &service.Service{DeliveryServiceApp: newMock}
			handler := controller.NewHandler(services)

			r := gin.New()
			r.GET("/deliveryservice/:id", handler.GetDeliveryServiceById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/deliveryservice/1", bytes.NewBufferString(tt.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())

		})
	}
}
func TestHandler_UpdateDeliveryService(t *testing.T) {
	type mockBehavior func(s *mock_service.MockDeliveryServiceApp, serv dao.DeliveryService)
	testTable := []struct {
		name               string
		inputBody          string
		inputService       dao.DeliveryService
		id                 int
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"name":"name","email":"email"}`,
			inputService: dao.DeliveryService{
				Id:    1,
				Name:  "name",
				Email: "email",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockDeliveryServiceApp, serv dao.DeliveryService) {
				s.EXPECT().UpdateDeliveryService(serv).Return(nil)
			},
			expectedStatusCode: 204,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			get := mock_service.NewMockDeliveryServiceApp(c)
			testCase.mockBehavior(get, testCase.inputService)
			services := &service.Service{DeliveryServiceApp: get}
			handler := controller.NewHandler(services)

			r := gin.New()
			r.PUT("/deliveryservice/:id", handler.UpdateDeliveryService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/deliveryservice/1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}

}
