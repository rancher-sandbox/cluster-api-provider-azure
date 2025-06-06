/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go
//
// Generated by this command:
//
//	mockgen -destination client_mock.go -package mock_virtualmachines -source ../client.go Client
//

// Package mock_virtualmachines is a generated GoMock package.
package mock_virtualmachines

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcompute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	gomock "go.uber.org/mock/gomock"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAsync mocks base method.
func (m *MockClient) CreateOrUpdateAsync(ctx context.Context, spec azure.ResourceSpecGetter, resumeToken string, parameters any) (any, *runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAsync", ctx, spec, resumeToken, parameters)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(*runtime.Poller[armcompute.VirtualMachinesClientCreateOrUpdateResponse])
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateOrUpdateAsync indicates an expected call of CreateOrUpdateAsync.
func (mr *MockClientMockRecorder) CreateOrUpdateAsync(ctx, spec, resumeToken, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAsync", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateAsync), ctx, spec, resumeToken, parameters)
}

// DeleteAsync mocks base method.
func (m *MockClient) DeleteAsync(ctx context.Context, spec azure.ResourceSpecGetter, resumeToken string) (*runtime.Poller[armcompute.VirtualMachinesClientDeleteResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsync", ctx, spec, resumeToken)
	ret0, _ := ret[0].(*runtime.Poller[armcompute.VirtualMachinesClientDeleteResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAsync indicates an expected call of DeleteAsync.
func (mr *MockClientMockRecorder) DeleteAsync(ctx, spec, resumeToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsync", reflect.TypeOf((*MockClient)(nil).DeleteAsync), ctx, spec, resumeToken)
}

// Get mocks base method.
func (m *MockClient) Get(arg0 context.Context, arg1 azure.ResourceSpecGetter) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0, arg1)
}
