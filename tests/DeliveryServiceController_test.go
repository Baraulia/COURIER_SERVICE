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
)

func TestHandler_CreateDeliveryService(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, service dao.DeliveryService)

	testTable := []struct {
		name                   string
		inputBody              string
		inputService           dao.DeliveryService
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
			inputBody: `{"name":"test", "email":"test", "photo":"test", "description": "test","phone_number":"1234567", "status": "active"}`,
			inputService: dao.DeliveryService{
				Name:        "test",
				Email:       "test",
				Photo:       "test",
				Description: "test",
				Status:      "active",
				PhoneNumber: "1234567",
			},
			inputPerms: "",
			inputRole:  "Courier manager",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},
			mockBehavior: func(s *mock_service.MockAllProjectApp, service dao.DeliveryService) {
				s.EXPECT().CreateDeliveryService(service).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:         "Invalid request",
			inputBody:    "a",
			inputService: dao.DeliveryService{},
			inputPerms:   "",
			inputRole:    "Courier manager",
			inputToken:   "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},
			mockBehavior:        func(r *mock_service.MockAllProjectApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Invalid request"}`,
		},
		{
			name:         "empty fields",
			inputBody:    "{}",
			inputService: dao.DeliveryService{},
			inputPerms:   "",
			inputRole:    "Courier manager",
			inputToken:   "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},
			mockBehavior:        func(r *mock_service.MockAllProjectApp, service dao.DeliveryService) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"empty fields"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(newMock, testCase.inputService)
			testCase.mockBehaviorParseToken(newMock, testCase.inputToken)
			testCase.mockBehaviorCheck(newMock, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: newMock}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/deliveryservice/", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}

func TestHandler_GetAllDeliveryServices(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, service dao.DeliveryService)

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
		name                   string
		inputBody              string
		inputService           dao.DeliveryService
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
			name:       "OK",
			inputPerms: "",
			inputRole:  "Courier manager",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},

			mockBehavior: func(s *mock_service.MockAllProjectApp, service dao.DeliveryService) {
				s.EXPECT().GetAllDeliveryServices().Return(servicess, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"id":1,"name":"name","email":"email","photo":"photo","description":"description","phone_number":"123","manager_id":1,"status":"active"}]}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(newMock, testCase.inputService)
			testCase.mockBehaviorParseToken(newMock, testCase.inputToken)
			testCase.mockBehaviorCheck(newMock, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: newMock}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/deliveryservice/", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}

func TestHandler_GetDeliveryServiceById(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, service *dao.DeliveryService)

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
		name                   string
		inputBody              string
		inputService           dao.DeliveryService
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
			inputService: dao.DeliveryService{
				Id: 1,
			},
			mockBehavior: func(s *mock_service.MockAllProjectApp, service *dao.DeliveryService) {
				s.EXPECT().GetDeliveryServiceById(1).Return(serv, nil)
			},
			inputPerms: "",
			inputRole:  "Courier manager",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"name":"name","email":"email","photo":"photo","description":"description","phone_number":"123","manager_id":1,"status":"active"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			newMock := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(newMock, &testCase.inputService)
			testCase.mockBehaviorParseToken(newMock, testCase.inputToken)
			testCase.mockBehaviorCheck(newMock, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: newMock}
			handler := controller.NewHandler(services)
			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/deliveryservice/1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}
func TestHandler_UpdateDeliveryService(t *testing.T) {
	type mockBehaviorCheck func(s *mock_service.MockAllProjectApp, perms, role string)
	type mockBehaviorParseToken func(s *mock_service.MockAllProjectApp, token string)
	type mockBehavior func(s *mock_service.MockAllProjectApp, serv dao.DeliveryService)

	testTable := []struct {
		name                   string
		inputBody              string
		inputService           dao.DeliveryService
		id                     int
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
			inputBody: `{"name":"name","email":"email"}`,
			inputService: dao.DeliveryService{
				Id:    1,
				Name:  "name",
				Email: "email",
			},
			id: 1,
			mockBehavior: func(s *mock_service.MockAllProjectApp, serv dao.DeliveryService) {
				s.EXPECT().UpdateDeliveryService(serv).Return(nil)
			},
			inputPerms: "",
			inputRole:  "Courier manager",
			inputToken: "testToken",
			mockBehaviorParseToken: func(s *mock_service.MockAllProjectApp, token string) {
				s.EXPECT().ParseToken(token).Return(&courierProto.UserRole{
					UserId:      1,
					Role:        "Courier manager",
					Permissions: "",
				}, nil)
			},
			mockBehaviorCheck: func(s *mock_service.MockAllProjectApp, perms, role string) {
				s.EXPECT().CheckRoleRights(nil, "Courier manager", perms, role).Return(nil)
			},
			expectedStatusCode: 204,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			get := mock_service.NewMockAllProjectApp(c)
			testCase.mockBehavior(get, testCase.inputService)
			testCase.mockBehaviorParseToken(get, testCase.inputToken)
			testCase.mockBehaviorCheck(get, testCase.inputPerms, testCase.inputRole)

			services := &service.Service{AllProjectApp: get}
			handler := controller.NewHandler(services)

			r := handler.InitRoutesGin()

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/deliveryservice/1", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Authorization", "Bearer testToken")
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}

}
