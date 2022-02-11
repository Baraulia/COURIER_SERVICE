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

func TestHandler_GetCouriers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCourierApp, courier db.SmallInfo)
	var couriers []db.SmallInfo
	cour := db.SmallInfo{
		IdCourier:   1,
		CourierName: "test",
		PhoneNumber: "1038812",
		Photo:       "my fav photo",
		Surname:     "Shorokhov",
	}
	couriers = append(couriers, cour)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        db.SmallInfo
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov"}`,
			inputCourier: db.SmallInfo{
				IdCourier:   1,
				CourierName: "test",
				PhoneNumber: "1038812",
				Photo:       "my fav photo",
				Surname:     "Shorokhov",
			},
			mockBehavior: func(s *mock_service.MockCourierApp, courier db.SmallInfo) {
				s.EXPECT().GetCouriers().Return(couriers, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockCourierApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{CourierApp: get}
			handler := Controllers.NewHandler(services)

			r := mux.NewRouter()

			r.HandleFunc("/couriers", handler.GetCouriers).Methods("GET")

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
	type mockBehavior func(s *mock_service.MockCourierApp, courier db.SmallInfo)
	var couriers []db.SmallInfo
	cour := db.SmallInfo{
		IdCourier:   1,
		CourierName: "test",
		PhoneNumber: "1038812",
		Photo:       "my fav photo",
		Surname:     "Shorokhov",
	}
	couriers = append(couriers, cour)

	testTable := []struct {
		name                string
		inputBody           string
		inputCourier        db.SmallInfo
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov"}`,
			inputCourier: db.SmallInfo{
				IdCourier:   1,
				CourierName: "test",
				PhoneNumber: "1038812",
				Photo:       "my fav photo",
				Surname:     "Shorokhov",
			},
			mockBehavior: func(s *mock_service.MockCourierApp, courier db.SmallInfo) {
				s.EXPECT().GetOneCourier(1).Return(couriers, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id_courier":1,"courier_name":"test","phone_number":"1038812","photo":"my fav photo","surname":"Shorokhov"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockCourierApp(c)
			testCase.mockBehavior(get, testCase.inputCourier)

			services := &service.Service{CourierApp: get}
			handler := Controllers.NewHandler(services)

			r := mux.NewRouter()

			r.HandleFunc("/courier", handler.GetOneCourier).Methods("GET")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/courier?id=1", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Contains(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
