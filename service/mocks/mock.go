// Code generated by MockGen. DO NOT EDIT.
// Source: Service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	dao "github.com/Baraulia/COURIER_SERVICE/dao"
	gomock "github.com/golang/mock/gomock"
)

// MockOrderApp is a mock of OrderApp interface.
type MockOrderApp struct {
	ctrl     *gomock.Controller
	recorder *MockOrderAppMockRecorder
}

// MockOrderAppMockRecorder is the mock recorder for MockOrderApp.
type MockOrderAppMockRecorder struct {
	mock *MockOrderApp
}

// NewMockOrderApp creates a new mock instance.
func NewMockOrderApp(ctrl *gomock.Controller) *MockOrderApp {
	mock := &MockOrderApp{ctrl: ctrl}
	mock.recorder = &MockOrderAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderApp) EXPECT() *MockOrderAppMockRecorder {
	return m.recorder
}

// AssigningOrderToCourier mocks base method.
func (m *MockOrderApp) AssigningOrderToCourier(order dao.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssigningOrderToCourier", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssigningOrderToCourier indicates an expected call of AssigningOrderToCourier.
func (mr *MockOrderAppMockRecorder) AssigningOrderToCourier(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssigningOrderToCourier", reflect.TypeOf((*MockOrderApp)(nil).AssigningOrderToCourier), order)
}

// ChangeOrderStatus mocks base method.
func (m *MockOrderApp) ChangeOrderStatus(id uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", id)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockOrderAppMockRecorder) ChangeOrderStatus(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockOrderApp)(nil).ChangeOrderStatus), id)
}

// GetAllOrdersOfCourierService mocks base method.
func (m *MockOrderApp) GetAllOrdersOfCourierService(limit, page, idService int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrdersOfCourierService", limit, page, idService)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrdersOfCourierService indicates an expected call of GetAllOrdersOfCourierService.
func (mr *MockOrderAppMockRecorder) GetAllOrdersOfCourierService(limit, page, idService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrdersOfCourierService", reflect.TypeOf((*MockOrderApp)(nil).GetAllOrdersOfCourierService), limit, page, idService)
}

// GetCourierCompletedOrders mocks base method.
func (m *MockOrderApp) GetCourierCompletedOrders(limit, page, idCourier int) ([]dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrders", limit, page, idCourier)
	ret0, _ := ret[0].([]dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrders indicates an expected call of GetCourierCompletedOrders.
func (mr *MockOrderAppMockRecorder) GetCourierCompletedOrders(limit, page, idCourier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrders", reflect.TypeOf((*MockOrderApp)(nil).GetCourierCompletedOrders), limit, page, idCourier)
}

// GetCourierCompletedOrdersByMonth mocks base method.
func (m *MockOrderApp) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourierCompletedOrdersByMonth", limit, page, idService, Month, Year)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourierCompletedOrdersByMonth indicates an expected call of GetCourierCompletedOrdersByMonth.
func (mr *MockOrderAppMockRecorder) GetCourierCompletedOrdersByMonth(limit, page, idService, Month, Year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourierCompletedOrdersByMonth", reflect.TypeOf((*MockOrderApp)(nil).GetCourierCompletedOrdersByMonth), limit, page, idService, Month, Year)
}

// GetDetailedOrderById mocks base method.
func (m *MockOrderApp) GetDetailedOrderById(Id int) (*dao.DetailedOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailedOrderById", Id)
	ret0, _ := ret[0].(*dao.DetailedOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailedOrderById indicates an expected call of GetDetailedOrderById.
func (mr *MockOrderAppMockRecorder) GetDetailedOrderById(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailedOrderById", reflect.TypeOf((*MockOrderApp)(nil).GetDetailedOrderById), Id)
}

// GetOrder mocks base method.
func (m *MockOrderApp) GetOrder(id int) (dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderAppMockRecorder) GetOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderApp)(nil).GetOrder), id)
}

// GetOrders mocks base method.
func (m *MockOrderApp) GetOrders(id int) ([]dao.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", id)
	ret0, _ := ret[0].([]dao.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderAppMockRecorder) GetOrders(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderApp)(nil).GetOrders), id)
}

// MockCourierApp is a mock of CourierApp interface.
type MockCourierApp struct {
	ctrl     *gomock.Controller
	recorder *MockCourierAppMockRecorder
}

// MockCourierAppMockRecorder is the mock recorder for MockCourierApp.
type MockCourierAppMockRecorder struct {
	mock *MockCourierApp
}

// NewMockCourierApp creates a new mock instance.
func NewMockCourierApp(ctrl *gomock.Controller) *MockCourierApp {
	mock := &MockCourierApp{ctrl: ctrl}
	mock.recorder = &MockCourierAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCourierApp) EXPECT() *MockCourierAppMockRecorder {
	return m.recorder
}

// GetCourier mocks base method.
func (m *MockCourierApp) GetCourier(id int) (dao.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourier", id)
	ret0, _ := ret[0].(dao.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourier indicates an expected call of GetCourier.
func (mr *MockCourierAppMockRecorder) GetCourier(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourier", reflect.TypeOf((*MockCourierApp)(nil).GetCourier), id)
}

// GetCouriers mocks base method.
func (m *MockCourierApp) GetCouriers() ([]dao.SmallInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCouriers")
	ret0, _ := ret[0].([]dao.SmallInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCouriers indicates an expected call of GetCouriers.
func (mr *MockCourierAppMockRecorder) GetCouriers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCouriers", reflect.TypeOf((*MockCourierApp)(nil).GetCouriers))
}

// SaveCourier mocks base method.
func (m *MockCourierApp) SaveCourier(courier *dao.Courier) (*dao.Courier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCourier", courier)
	ret0, _ := ret[0].(*dao.Courier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCourier indicates an expected call of SaveCourier.
func (mr *MockCourierAppMockRecorder) SaveCourier(courier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCourier", reflect.TypeOf((*MockCourierApp)(nil).SaveCourier), courier)
}

// MockDeliveryServiceApp is a mock of DeliveryServiceApp interface.
type MockDeliveryServiceApp struct {
	ctrl     *gomock.Controller
	recorder *MockDeliveryServiceAppMockRecorder
}

// MockDeliveryServiceAppMockRecorder is the mock recorder for MockDeliveryServiceApp.
type MockDeliveryServiceAppMockRecorder struct {
	mock *MockDeliveryServiceApp
}

// NewMockDeliveryServiceApp creates a new mock instance.
func NewMockDeliveryServiceApp(ctrl *gomock.Controller) *MockDeliveryServiceApp {
	mock := &MockDeliveryServiceApp{ctrl: ctrl}
	mock.recorder = &MockDeliveryServiceAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeliveryServiceApp) EXPECT() *MockDeliveryServiceAppMockRecorder {
	return m.recorder
}

// CreateDeliveryService mocks base method.
func (m *MockDeliveryServiceApp) CreateDeliveryService(DeliveryService dao.DeliveryService) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeliveryService", DeliveryService)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeliveryService indicates an expected call of CreateDeliveryService.
func (mr *MockDeliveryServiceAppMockRecorder) CreateDeliveryService(DeliveryService interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryService", reflect.TypeOf((*MockDeliveryServiceApp)(nil).CreateDeliveryService), DeliveryService)
}
