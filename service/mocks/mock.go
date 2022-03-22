// Code generated by MockGen. DO NOT EDIT.
// Source: Service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPC"
	authProto "github.com/Baraulia/COURIER_SERVICE/GRPCC"
	dao "github.com/Baraulia/COURIER_SERVICE/dao"
	gomock "github.com/golang/mock/gomock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockAllProjectApp is a mock of AllProjectApp interface.
type MockAllProjectApp struct {
	ctrl     *gomock.Controller
	recorder *MockAllProjectAppMockRecorder
}

// MockAllProjectAppMockRecorder is the mock recorder for MockAllProjectApp.
type MockAllProjectAppMockRecorder struct {
	mock *MockAllProjectApp
}

// NewMockAllProjectApp creates a new mock instance.
func NewMockAllProjectApp(ctrl *gomock.Controller) *MockAllProjectApp {
	mock := &MockAllProjectApp{ctrl: ctrl}
	mock.recorder = &MockAllProjectAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllProjectApp) EXPECT() *MockAllProjectAppMockRecorder {
	return m.recorder
}

// AssigningOrderToCourier mocks base method.
func (m *MockAllProjectApp) AssigningOrderToCourier(order dao.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssigningOrderToCourier", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssigningOrderToCourier indicates an expected call of AssigningOrderToCourier.
func (mr *MockAllProjectAppMockRecorder) AssigningOrderToCourier(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssigningOrderToCourier", reflect.TypeOf((*MockAllProjectApp)(nil).AssigningOrderToCourier), order)
}

// ChangeOrderStatus mocks base method.
func (m *MockAllProjectApp) ChangeOrderStatus(text string, id uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", text, id)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockAllProjectAppMockRecorder) ChangeOrderStatus(text, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockAllProjectApp)(nil).ChangeOrderStatus), text, id)
}

// CheckRights mocks base method.
func (m *MockAllProjectApp) CheckRights(neededPerms []string, givenPerms string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRights", neededPerms, givenPerms)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRights indicates an expected call of CheckRights.
func (mr *MockAllProjectAppMockRecorder) CheckRights(neededPerms, givenPerms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRights", reflect.TypeOf((*MockAllProjectApp)(nil).CheckRights), neededPerms, givenPerms)
}

// CheckRole mocks base method.
func (m *MockAllProjectApp) CheckRole(neededRoles []string, givenRole string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRole", neededRoles, givenRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRole indicates an expected call of CheckRole.
func (mr *MockAllProjectAppMockRecorder) CheckRole(neededRoles, givenRole interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRole", reflect.TypeOf((*MockAllProjectApp)(nil).CheckRole), neededRoles, givenRole)
}

// CreateDeliveryService mocks base method.
func (m *MockAllProjectApp) CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeliveryService", DeliveryService)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeliveryService indicates an expected call of CreateDeliveryService.
func (mr *MockAllProjectAppMockRecorder) CreateDeliveryService(DeliveryService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryService", reflect.TypeOf((*MockAllProjectApp)(nil).CreateDeliveryService), DeliveryService)
}

// CreateOrder mocks base method.
func (m *MockAllProjectApp) CreateOrder(order *courierProto.OrderCourierServer) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockAllProjectAppMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockAllProjectApp)(nil).CreateOrder), order)
}

// GetAllDeliveryServices mocks base method.
func (m *MockAllProjectApp) GetAllDeliveryServices() ([]dao.DeliveryService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDeliveryServices")
	ret0, _ := ret[0].([]dao.DeliveryService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDeliveryServices indicates an expected call of GetAllDeliveryServices.
func (mr *MockAllProjectAppMockRecorder) GetAllDeliveryServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDeliveryServices", reflect.TypeOf((*MockAllProjectApp)(nil).GetAllDeliveryServices))
}

// GetAllOrdersOfCourierService mocks base method.
func (m *MockAllProjectApp) GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrdersOfCourierService", limit, page, idService)
	ret0, _ := ret[0].([]dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrdersOfCourierService indicates an expected call of GetAllOrdersOfCourierService.
func (mr *MockAllProjectAppMockRecorder) GetAllOrdersOfCourierService(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrdersOfCourierService", reflect.TypeOf((*MockAllProjectApp)(nil).GetAllOrdersOfCourierService), limit, page, idService)
}

// GetCompletedOrdersOfCourierService mocks base method.
func (m *MockAllProjectApp) GetCompletedOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletedOrdersOfCourierService", limit, page, idService)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompletedOrdersOfCourierService indicates an expected call of GetCompletedOrdersOfCourierService.
func (mr *MockAllProjectAppMockRecorder) GetCompletedOrdersOfCourierService(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletedOrdersOfCourierService", reflect.TypeOf((*MockAllProjectApp)(nil).GetCompletedOrdersOfCourierService), limit, page, idService)
}

// GetCompletedOrdersOfCourierServiceByCourierId mocks base method.
func (m *MockAllProjectApp) GetCompletedOrdersOfCourierServiceByCourierId(limit, page, idService int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletedOrdersOfCourierServiceByCourierId", limit, page, idService)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompletedOrdersOfCourierServiceByCourierId indicates an expected call of GetCompletedOrdersOfCourierServiceByCourierId.
func (mr *MockAllProjectAppMockRecorder) GetCompletedOrdersOfCourierServiceByCourierId(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletedOrdersOfCourierServiceByCourierId", reflect.TypeOf((*MockAllProjectApp)(nil).GetCompletedOrdersOfCourierServiceByCourierId), limit, page, idService)
}

// GetCompletedOrdersOfCourierServiceByDate mocks base method.
func (m *MockAllProjectApp) GetCompletedOrdersOfCourierServiceByDate(limit, page, idService int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletedOrdersOfCourierServiceByDate", limit, page, idService)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompletedOrdersOfCourierServiceByDate indicates an expected call of GetCompletedOrdersOfCourierServiceByDate.
func (mr *MockAllProjectAppMockRecorder) GetCompletedOrdersOfCourierServiceByDate(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletedOrdersOfCourierServiceByDate", reflect.TypeOf((*MockAllProjectApp)(nil).GetCompletedOrdersOfCourierServiceByDate), limit, page, idService)
}

// GetCourier mocks base method.
func (m *MockAllProjectApp) GetCourier(id int) (dao.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourier", id)
	ret0, _ := ret[0].(dao.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourier indicates an expected call of GetCourier.
func (mr *MockAllProjectAppMockRecorder) GetCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourier", reflect.TypeOf((*MockAllProjectApp)(nil).GetCourier), id)
}

// GetCourierCompletedOrders mocks base method.
func (m *MockAllProjectApp) GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrders", limit, page, idCourier)
	ret0, _ := ret[0].([]dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrders indicates an expected call of GetCourierCompletedOrders.
func (mr *MockAllProjectAppMockRecorder) GetCourierCompletedOrders(limit, page, idCourier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrders", reflect.TypeOf((*MockAllProjectApp)(nil).GetCourierCompletedOrders), limit, page, idCourier)
}

// GetCourierCompletedOrdersByMonth mocks base method.
func (m *MockAllProjectApp) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrdersByMonth", limit, page, idService, Month, Year)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrdersByMonth indicates an expected call of GetCourierCompletedOrdersByMonth.
func (mr *MockAllProjectAppMockRecorder) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrdersByMonth", reflect.TypeOf((*MockAllProjectApp)(nil).GetCourierCompletedOrdersByMonth), limit, page, idService, Month, Year)
}

// GetCouriers mocks base method.
func (m *MockAllProjectApp) GetCouriers() ([]dao.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCouriers")
	ret0, _ := ret[0].([]dao.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCouriers indicates an expected call of GetCouriers.
func (mr *MockAllProjectAppMockRecorder) GetCouriers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCouriers", reflect.TypeOf((*MockAllProjectApp)(nil).GetCouriers))
}

// GetDeliveryServiceById mocks base method.
func (m *MockAllProjectApp) GetDeliveryServiceById(Id int) (*dao.DeliveryService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeliveryServiceById", Id)
	ret0, _ := ret[0].(*dao.DeliveryService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeliveryServiceById indicates an expected call of GetDeliveryServiceById.
func (mr *MockAllProjectAppMockRecorder) GetDeliveryServiceById(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeliveryServiceById", reflect.TypeOf((*MockAllProjectApp)(nil).GetDeliveryServiceById), Id)
}

// GetDetailedOrderById mocks base method.
func (m *MockAllProjectApp) GetDetailedOrderById(Id int) (*dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailedOrderById", Id)
	ret0, _ := ret[0].(*dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailedOrderById indicates an expected call of GetDetailedOrderById.
func (mr *MockAllProjectAppMockRecorder) GetDetailedOrderById(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailedOrderById", reflect.TypeOf((*MockAllProjectApp)(nil).GetDetailedOrderById), Id)
}

// GetOrder mocks base method.
func (m *MockAllProjectApp) GetOrder(id int) (dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockAllProjectAppMockRecorder) GetOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockAllProjectApp)(nil).GetOrder), id)
}

// GetOrderForChange mocks base method.
func (m *MockAllProjectApp) GetOrderForChange(id int) (dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderForChange", id)
	ret0, _ := ret[0].(dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderForChange indicates an expected call of GetOrderForChange.
func (mr *MockAllProjectAppMockRecorder) GetOrderForChange(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderForChange", reflect.TypeOf((*MockAllProjectApp)(nil).GetOrderForChange), id)
}

// GetOrders mocks base method.
func (m *MockAllProjectApp) GetOrders(id int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", id)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockAllProjectAppMockRecorder) GetOrders(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockAllProjectApp)(nil).GetOrders), id)
}

// GetOrdersOfCourierServiceForManager mocks base method.
func (m *MockAllProjectApp) GetOrdersOfCourierServiceForManager(limit, page, idService int) ([]dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersOfCourierServiceForManager", limit, page, idService)
	ret0, _ := ret[0].([]dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersOfCourierServiceForManager indicates an expected call of GetOrdersOfCourierServiceForManager.
func (mr *MockAllProjectAppMockRecorder) GetOrdersOfCourierServiceForManager(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersOfCourierServiceForManager", reflect.TypeOf((*MockAllProjectApp)(nil).GetOrdersOfCourierServiceForManager), limit, page, idService)
}

// GetServices mocks base method.
func (m *MockAllProjectApp) GetServices(in *emptypb.Empty) (*courierProto.ServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServices", in)
	ret0, _ := ret[0].(*courierProto.ServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices.
func (mr *MockAllProjectAppMockRecorder) GetServices(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockAllProjectApp)(nil).GetServices), in)
}

// ParseToken mocks base method.
func (m *MockAllProjectApp) ParseToken(token string) (*authProto.UserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(*authProto.UserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAllProjectAppMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAllProjectApp)(nil).ParseToken), token)
}

// SaveCourier mocks base method.
func (m *MockAllProjectApp) SaveCourier(courier *dao.Courier) (*dao.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCourier", courier)
	ret0, _ := ret[0].(*dao.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCourier indicates an expected call of SaveCourier.
func (mr *MockAllProjectAppMockRecorder) SaveCourier(courier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCourier", reflect.TypeOf((*MockAllProjectApp)(nil).SaveCourier), courier)
}

// SaveCourierPhoto mocks base method.
func (m *MockAllProjectApp) SaveCourierPhoto(cover []byte, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCourierPhoto", cover, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCourierPhoto indicates an expected call of SaveCourierPhoto.
func (mr *MockAllProjectAppMockRecorder) SaveCourierPhoto(cover, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCourierPhoto", reflect.TypeOf((*MockAllProjectApp)(nil).SaveCourierPhoto), cover, id)
}

// SaveLogoFile mocks base method.
func (m *MockAllProjectApp) SaveLogoFile(cover []byte, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLogoFile", cover, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLogoFile indicates an expected call of SaveLogoFile.
func (mr *MockAllProjectAppMockRecorder) SaveLogoFile(cover, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLogoFile", reflect.TypeOf((*MockAllProjectApp)(nil).SaveLogoFile), cover, id)
}

// UpdateCourier mocks base method.
func (m *MockAllProjectApp) UpdateCourier(id uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourier", id)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourier indicates an expected call of UpdateCourier.
func (mr *MockAllProjectAppMockRecorder) UpdateCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourier", reflect.TypeOf((*MockAllProjectApp)(nil).UpdateCourier), id)
}

// UpdateDeliveryService mocks base method.
func (m *MockAllProjectApp) UpdateDeliveryService(service dao.DeliveryService) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeliveryService", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeliveryService indicates an expected call of UpdateDeliveryService.
func (mr *MockAllProjectAppMockRecorder) UpdateDeliveryService(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeliveryService", reflect.TypeOf((*MockAllProjectApp)(nil).UpdateDeliveryService), service)
}
