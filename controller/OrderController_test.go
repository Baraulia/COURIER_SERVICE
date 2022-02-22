package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"stlab.itechart-group.com/go/food_delivery/courier_service/model"
	mocks "stlab.itechart-group.com/go/food_delivery/courier_service/model/mocks"
	"testing"
)

func TestHandler_AssigningOrderToCourier(t *testing.T) {
	type mockBehavior func(s *mocks.MockOrderApp, order dao.Order)
	testTable := []struct {
		name               string
		inputBody          string
		inputOrder         dao.Order
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"id":15, "courier_id":8}`,
			inputOrder: dao.Order{
				Id:   15,
				IdCourier: 8,
			},
			mockBehavior: func(s *mocks.MockOrderApp, order dao.Order) {
				s.EXPECT().AssigningOrderToCourier(order).Return(nil)
			},
			expectedStatusCode: 200,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mocks.NewMockOrderApp(c)
			testCase.mockBehavior(get, testCase.inputOrder)

			services := &model.Service{OrderApp: get}
			handler := NewHandler(services)

			r := gin.New()

			r.PUT("/orders", handler.UpdateOrder)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/orders", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}

