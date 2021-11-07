// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/template-go-server/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/template-go-server/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAuthor mocks base method.
func (m *MockStore) CreateAuthor(arg0 context.Context, arg1 db.CreateAuthorParams) (db.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthor", arg0, arg1)
	ret0, _ := ret[0].(db.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthor indicates an expected call of CreateAuthor.
func (mr *MockStoreMockRecorder) CreateAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthor", reflect.TypeOf((*MockStore)(nil).CreateAuthor), arg0, arg1)
}

// DeleteAuthor mocks base method.
func (m *MockStore) DeleteAuthor(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAuthor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAuthor indicates an expected call of DeleteAuthor.
func (mr *MockStoreMockRecorder) DeleteAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAuthor", reflect.TypeOf((*MockStore)(nil).DeleteAuthor), arg0, arg1)
}

// GetAuthor mocks base method.
func (m *MockStore) GetAuthor(arg0 context.Context, arg1 int32) (db.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthor", arg0, arg1)
	ret0, _ := ret[0].(db.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthor indicates an expected call of GetAuthor.
func (mr *MockStoreMockRecorder) GetAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthor", reflect.TypeOf((*MockStore)(nil).GetAuthor), arg0, arg1)
}

// ListAuthors mocks base method.
func (m *MockStore) ListAuthors(arg0 context.Context, arg1 db.ListAuthorsParams) ([]db.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthors", arg0, arg1)
	ret0, _ := ret[0].([]db.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthors indicates an expected call of ListAuthors.
func (mr *MockStoreMockRecorder) ListAuthors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthors", reflect.TypeOf((*MockStore)(nil).ListAuthors), arg0, arg1)
}

// UpdateAuthor mocks base method.
func (m *MockStore) UpdateAuthor(arg0 context.Context, arg1 db.UpdateAuthorParams) (db.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthor", arg0, arg1)
	ret0, _ := ret[0].(db.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthor indicates an expected call of UpdateAuthor.
func (mr *MockStoreMockRecorder) UpdateAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthor", reflect.TypeOf((*MockStore)(nil).UpdateAuthor), arg0, arg1)
}