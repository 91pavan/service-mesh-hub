// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_zephyr_core is a generated GoMock package.
package mock_zephyr_core

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockMeshClient is a mock of MeshClient interface
type MockMeshClient struct {
	ctrl     *gomock.Controller
	recorder *MockMeshClientMockRecorder
}

// MockMeshClientMockRecorder is the mock recorder for MockMeshClient
type MockMeshClientMockRecorder struct {
	mock *MockMeshClient
}

// NewMockMeshClient creates a new mock instance
func NewMockMeshClient(ctrl *gomock.Controller) *MockMeshClient {
	mock := &MockMeshClient{ctrl: ctrl}
	mock.recorder = &MockMeshClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshClient) EXPECT() *MockMeshClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMeshClient) Create(ctx context.Context, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMeshClientMockRecorder) Create(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMeshClient)(nil).Create), ctx, mesh)
}

// Delete mocks base method
func (m *MockMeshClient) Delete(ctx context.Context, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMeshClientMockRecorder) Delete(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshClient)(nil).Delete), ctx, mesh)
}

// Get mocks base method
func (m *MockMeshClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMeshClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMeshClient)(nil).Get), ctx, objKey)
}

// List mocks base method
func (m *MockMeshClient) List(ctx context.Context, opts ...client.ListOption) (*v1alpha1.MeshList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1alpha1.MeshList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockMeshClientMockRecorder) List(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMeshClient)(nil).List), varargs...)
}

// MockMeshWorkloadClient is a mock of MeshWorkloadClient interface
type MockMeshWorkloadClient struct {
	ctrl     *gomock.Controller
	recorder *MockMeshWorkloadClientMockRecorder
}

// MockMeshWorkloadClientMockRecorder is the mock recorder for MockMeshWorkloadClient
type MockMeshWorkloadClientMockRecorder struct {
	mock *MockMeshWorkloadClient
}

// NewMockMeshWorkloadClient creates a new mock instance
func NewMockMeshWorkloadClient(ctrl *gomock.Controller) *MockMeshWorkloadClient {
	mock := &MockMeshWorkloadClient{ctrl: ctrl}
	mock.recorder = &MockMeshWorkloadClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshWorkloadClient) EXPECT() *MockMeshWorkloadClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMeshWorkloadClient) Create(ctx context.Context, mesh *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMeshWorkloadClientMockRecorder) Create(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMeshWorkloadClient)(nil).Create), ctx, mesh)
}

// Delete mocks base method
func (m *MockMeshWorkloadClient) Delete(ctx context.Context, mesh *v1alpha1.MeshWorkload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMeshWorkloadClientMockRecorder) Delete(ctx, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMeshWorkloadClient)(nil).Delete), ctx, mesh)
}

// Get mocks base method
func (m *MockMeshWorkloadClient) Get(ctx context.Context, objKey client.ObjectKey) (*v1alpha1.MeshWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, objKey)
	ret0, _ := ret[0].(*v1alpha1.MeshWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMeshWorkloadClientMockRecorder) Get(ctx, objKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMeshWorkloadClient)(nil).Get), ctx, objKey)
}
