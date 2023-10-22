// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	go_genderize "github.com/SteelPangolin/go-genderize"
	domain "github.com/danluki/effective-mobile-test/internal/domain"
	service "github.com/danluki/effective-mobile-test/internal/service"
	gomock "github.com/golang/mock/gomock"
	agify "github.com/masonkmeyer/agify"
	nationalize "github.com/masonkmeyer/nationalize"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsers) Create(ctx context.Context, input service.CreateUserInput) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockUsers) Delete(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsers)(nil).Delete), ctx, id)
}

// List mocks base method.
func (m *MockUsers) List(ctx context.Context, input service.ListUsersInput) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, input)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsersMockRecorder) List(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsers)(nil).List), ctx, input)
}

// Update mocks base method.
func (m *MockUsers) Update(ctx context.Context, input service.UpdateUserInput) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUsersMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsers)(nil).Update), ctx, input)
}

// MockGenderizeClientInterface is a mock of GenderizeClientInterface interface.
type MockGenderizeClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGenderizeClientInterfaceMockRecorder
}

// MockGenderizeClientInterfaceMockRecorder is the mock recorder for MockGenderizeClientInterface.
type MockGenderizeClientInterfaceMockRecorder struct {
	mock *MockGenderizeClientInterface
}

// NewMockGenderizeClientInterface creates a new mock instance.
func NewMockGenderizeClientInterface(ctrl *gomock.Controller) *MockGenderizeClientInterface {
	mock := &MockGenderizeClientInterface{ctrl: ctrl}
	mock.recorder = &MockGenderizeClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenderizeClientInterface) EXPECT() *MockGenderizeClientInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGenderizeClientInterface) Get(query go_genderize.Query) ([]go_genderize.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", query)
	ret0, _ := ret[0].([]go_genderize.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGenderizeClientInterfaceMockRecorder) Get(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGenderizeClientInterface)(nil).Get), query)
}

// MockNationalizeClientInterface is a mock of NationalizeClientInterface interface.
type MockNationalizeClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNationalizeClientInterfaceMockRecorder
}

// MockNationalizeClientInterfaceMockRecorder is the mock recorder for MockNationalizeClientInterface.
type MockNationalizeClientInterfaceMockRecorder struct {
	mock *MockNationalizeClientInterface
}

// NewMockNationalizeClientInterface creates a new mock instance.
func NewMockNationalizeClientInterface(ctrl *gomock.Controller) *MockNationalizeClientInterface {
	mock := &MockNationalizeClientInterface{ctrl: ctrl}
	mock.recorder = &MockNationalizeClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNationalizeClientInterface) EXPECT() *MockNationalizeClientInterfaceMockRecorder {
	return m.recorder
}

// Predict mocks base method.
func (m *MockNationalizeClientInterface) Predict(name string) (*nationalize.Prediction, *nationalize.RateLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Predict", name)
	ret0, _ := ret[0].(*nationalize.Prediction)
	ret1, _ := ret[1].(*nationalize.RateLimit)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Predict indicates an expected call of Predict.
func (mr *MockNationalizeClientInterfaceMockRecorder) Predict(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Predict", reflect.TypeOf((*MockNationalizeClientInterface)(nil).Predict), name)
}

// MockAgifyClientInterface is a mock of AgifyClientInterface interface.
type MockAgifyClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAgifyClientInterfaceMockRecorder
}

// MockAgifyClientInterfaceMockRecorder is the mock recorder for MockAgifyClientInterface.
type MockAgifyClientInterfaceMockRecorder struct {
	mock *MockAgifyClientInterface
}

// NewMockAgifyClientInterface creates a new mock instance.
func NewMockAgifyClientInterface(ctrl *gomock.Controller) *MockAgifyClientInterface {
	mock := &MockAgifyClientInterface{ctrl: ctrl}
	mock.recorder = &MockAgifyClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgifyClientInterface) EXPECT() *MockAgifyClientInterfaceMockRecorder {
	return m.recorder
}

// Predict mocks base method.
func (m *MockAgifyClientInterface) Predict(name string) (*agify.Prediction, *agify.RateLimit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Predict", name)
	ret0, _ := ret[0].(*agify.Prediction)
	ret1, _ := ret[1].(*agify.RateLimit)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Predict indicates an expected call of Predict.
func (mr *MockAgifyClientInterfaceMockRecorder) Predict(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Predict", reflect.TypeOf((*MockAgifyClientInterface)(nil).Predict), name)
}