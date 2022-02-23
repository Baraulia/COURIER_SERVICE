package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"stlab.itechart-group.com/go/food_delivery/courier_service/model"
	mocks "stlab.itechart-group.com/go/food_delivery/courier_service/model/mocks"
	"testing"
)

func TestHandler_CreateDeliveryService(t *testing.T) {
	type mockBehavior func(s *mocks.MockDeliveryServiceApp, service dao.DeliveryService)
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
			inputBody: `{"name":"test", "email":"test", "photo":"test", "working_now":true, "description": "test", "deleted": false}`,
			inputService: dao.DeliveryService{
				Name:        "test",
				Email:       "test",
				Photo:       "test",
				WorkingNow:  true,
				Description: "test",
				Deleted:     false,
			},
			mockBehavior: func(s *mocks.MockDeliveryServiceApp, service dao.DeliveryService) {
				s.EXPECT().CreateDeliveryService(service).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Invalid request",
			inputBody:           `{"email": "email",}`,
			inputService:        dao.DeliveryService{},
			mockBehavior:        func(r *mocks.MockDeliveryServiceApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
		{
			name:                "empty fields",
			inputBody:           `{"email": "email"}`,
			inputService:        dao.DeliveryService{},
			mockBehavior:        func(r *mocks.MockDeliveryServiceApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"empty fields"}`,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mocks.NewMockDeliveryServiceApp(c)
			tt.mockBehavior(newMock, tt.inputService)
			services := &model.Service{DeliveryServiceApp: newMock}
			handler := NewHandler(services)
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
