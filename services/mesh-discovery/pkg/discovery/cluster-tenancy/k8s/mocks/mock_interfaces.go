// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_k8s_tenancy is a generated GoMock package.
package mock_k8s_tenancy

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	controller0 "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
	v1 "k8s.io/api/core/v1"
)

// MockClusterTenancyRegistrarLoop is a mock of ClusterTenancyRegistrarLoop interface.
type MockClusterTenancyRegistrarLoop struct {
	ctrl     *gomock.Controller
	recorder *MockClusterTenancyRegistrarLoopMockRecorder
}

// MockClusterTenancyRegistrarLoopMockRecorder is the mock recorder for MockClusterTenancyRegistrarLoop.
type MockClusterTenancyRegistrarLoopMockRecorder struct {
	mock *MockClusterTenancyRegistrarLoop
}

// NewMockClusterTenancyRegistrarLoop creates a new mock instance.
func NewMockClusterTenancyRegistrarLoop(ctrl *gomock.Controller) *MockClusterTenancyRegistrarLoop {
	mock := &MockClusterTenancyRegistrarLoop{ctrl: ctrl}
	mock.recorder = &MockClusterTenancyRegistrarLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterTenancyRegistrarLoop) EXPECT() *MockClusterTenancyRegistrarLoopMockRecorder {
	return m.recorder
}

// StartRegistration mocks base method.
func (m *MockClusterTenancyRegistrarLoop) StartRegistration(ctx context.Context, podEventWatcher controller0.PodEventWatcher, meshEventWatcher controller.MeshEventWatcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartRegistration", ctx, podEventWatcher, meshEventWatcher)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartRegistration indicates an expected call of StartRegistration.
func (mr *MockClusterTenancyRegistrarLoopMockRecorder) StartRegistration(ctx, podEventWatcher, meshEventWatcher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRegistration", reflect.TypeOf((*MockClusterTenancyRegistrarLoop)(nil).StartRegistration), ctx, podEventWatcher, meshEventWatcher)
}

// MockClusterTenancyRegistrar is a mock of ClusterTenancyRegistrar interface.
type MockClusterTenancyRegistrar struct {
	ctrl     *gomock.Controller
	recorder *MockClusterTenancyRegistrarMockRecorder
}

// MockClusterTenancyRegistrarMockRecorder is the mock recorder for MockClusterTenancyRegistrar.
type MockClusterTenancyRegistrarMockRecorder struct {
	mock *MockClusterTenancyRegistrar
}

// NewMockClusterTenancyRegistrar creates a new mock instance.
func NewMockClusterTenancyRegistrar(ctrl *gomock.Controller) *MockClusterTenancyRegistrar {
	mock := &MockClusterTenancyRegistrar{ctrl: ctrl}
	mock.recorder = &MockClusterTenancyRegistrarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterTenancyRegistrar) EXPECT() *MockClusterTenancyRegistrarMockRecorder {
	return m.recorder
}

// MeshFromSidecar mocks base method.
func (m *MockClusterTenancyRegistrar) MeshFromSidecar(ctx context.Context, pod *v1.Pod) (*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MeshFromSidecar", ctx, pod)
	ret0, _ := ret[0].(*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MeshFromSidecar indicates an expected call of MeshFromSidecar.
func (mr *MockClusterTenancyRegistrarMockRecorder) MeshFromSidecar(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MeshFromSidecar", reflect.TypeOf((*MockClusterTenancyRegistrar)(nil).MeshFromSidecar), ctx, pod)
}

// RegisterMesh mocks base method.
func (m *MockClusterTenancyRegistrar) RegisterMesh(ctx context.Context, clusterName string, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMesh", ctx, clusterName, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterMesh indicates an expected call of RegisterMesh.
func (mr *MockClusterTenancyRegistrarMockRecorder) RegisterMesh(ctx, clusterName, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMesh", reflect.TypeOf((*MockClusterTenancyRegistrar)(nil).RegisterMesh), ctx, clusterName, mesh)
}

// DeregisterMesh mocks base method.
func (m *MockClusterTenancyRegistrar) DeregisterMesh(ctx context.Context, clusterName string, mesh *v1alpha1.Mesh) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterMesh", ctx, clusterName, mesh)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterMesh indicates an expected call of DeregisterMesh.
func (mr *MockClusterTenancyRegistrarMockRecorder) DeregisterMesh(ctx, clusterName, mesh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterMesh", reflect.TypeOf((*MockClusterTenancyRegistrar)(nil).DeregisterMesh), ctx, clusterName, mesh)
}
