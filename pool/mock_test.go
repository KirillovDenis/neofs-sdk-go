// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nspcc-dev/neofs-sdk-go/pool (interfaces: Client)

// Package pool is a generated GoMock package.
package pool

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/nspcc-dev/neofs-api-go/v2/rpc/client"
	client0 "github.com/nspcc-dev/neofs-sdk-go/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
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

// AnnounceContainerUsedSpace mocks base method.
func (m *MockClient) AnnounceContainerUsedSpace(arg0 context.Context, arg1 client0.AnnounceSpacePrm) (*client0.AnnounceSpaceRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "AnnounceContainerUsedSpace", varargs...)
	ret0, _ := ret[0].(*client0.AnnounceSpaceRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceContainerUsedSpace indicates an expected call of AnnounceContainerUsedSpace.
func (mr *MockClientMockRecorder) AnnounceContainerUsedSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceContainerUsedSpace", reflect.TypeOf((*MockClient)(nil).AnnounceContainerUsedSpace), varargs...)
}

// AnnounceIntermediateTrust mocks base method.
func (m *MockClient) AnnounceIntermediateTrust(arg0 context.Context, arg1 client0.AnnounceIntermediateTrustPrm) (*client0.AnnounceIntermediateTrustRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "AnnounceIntermediateTrust", varargs...)
	ret0, _ := ret[0].(*client0.AnnounceIntermediateTrustRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceIntermediateTrust indicates an expected call of AnnounceIntermediateTrust.
func (mr *MockClientMockRecorder) AnnounceIntermediateTrust(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceIntermediateTrust", reflect.TypeOf((*MockClient)(nil).AnnounceIntermediateTrust), varargs...)
}

// AnnounceLocalTrust mocks base method.
func (m *MockClient) AnnounceLocalTrust(arg0 context.Context, arg1 client0.AnnounceLocalTrustPrm) (*client0.AnnounceLocalTrustRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "AnnounceLocalTrust", varargs...)
	ret0, _ := ret[0].(*client0.AnnounceLocalTrustRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceLocalTrust indicates an expected call of AnnounceLocalTrust.
func (mr *MockClientMockRecorder) AnnounceLocalTrust(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceLocalTrust", reflect.TypeOf((*MockClient)(nil).AnnounceLocalTrust), varargs...)
}

// Conn mocks base method.
func (m *MockClient) Conn() io.Closer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(io.Closer)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockClientMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockClient)(nil).Conn))
}

// CreateSession mocks base method.
func (m *MockClient) CreateSession(arg0 context.Context, arg1 client0.CreateSessionPrm) (*client0.CreateSessionRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "CreateSession", varargs...)
	ret0, _ := ret[0].(*client0.CreateSessionRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockClientMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockClient)(nil).CreateSession), varargs...)
}

// DeleteContainer mocks base method.
func (m *MockClient) DeleteContainer(arg0 context.Context, arg1 client0.ContainerDeletePrm) (*client0.ContainerDeleteRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "DeleteContainer", varargs...)
	ret0, _ := ret[0].(*client0.ContainerDeleteRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteContainer indicates an expected call of DeleteContainer.
func (mr *MockClientMockRecorder) DeleteContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockClient)(nil).DeleteContainer), varargs...)
}

// DeleteObject mocks base method.
func (m *MockClient) DeleteObject(arg0 context.Context, arg1 *client0.DeleteObjectParams, arg2 ...client0.CallOption) (*client0.ObjectDeleteRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObject", varargs...)
	ret0, _ := ret[0].(*client0.ObjectDeleteRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockClientMockRecorder) DeleteObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockClient)(nil).DeleteObject), varargs...)
}

// EACL mocks base method.
func (m *MockClient) EACL(arg0 context.Context, arg1 client0.EACLPrm) (*client0.EACLRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "EACL", varargs...)
	ret0, _ := ret[0].(*client0.EACLRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EACL indicates an expected call of EACL.
func (mr *MockClientMockRecorder) EACL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EACL", reflect.TypeOf((*MockClient)(nil).EACL), varargs...)
}

// EndpointInfo mocks base method.
func (m *MockClient) EndpointInfo(arg0 context.Context, arg1 client0.EndpointInfoPrm) (*client0.EndpointInfoRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "EndpointInfo", varargs...)
	ret0, _ := ret[0].(*client0.EndpointInfoRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointInfo indicates an expected call of EndpointInfo.
func (mr *MockClientMockRecorder) EndpointInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointInfo", reflect.TypeOf((*MockClient)(nil).EndpointInfo), varargs...)
}

// GetBalance mocks base method.
func (m *MockClient) GetBalance(arg0 context.Context, arg1 client0.GetBalancePrm) (*client0.GetBalanceRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "GetBalance", varargs...)
	ret0, _ := ret[0].(*client0.GetBalanceRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockClientMockRecorder) GetBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockClient)(nil).GetBalance), varargs...)
}

// GetContainer mocks base method.
func (m *MockClient) GetContainer(arg0 context.Context, arg1 client0.ContainerGetPrm) (*client0.ContainerGetRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "GetContainer", varargs...)
	ret0, _ := ret[0].(*client0.ContainerGetRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainer indicates an expected call of GetContainer.
func (mr *MockClientMockRecorder) GetContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainer", reflect.TypeOf((*MockClient)(nil).GetContainer), varargs...)
}

// GetObject mocks base method.
func (m *MockClient) GetObject(arg0 context.Context, arg1 *client0.GetObjectParams, arg2 ...client0.CallOption) (*client0.ObjectGetRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*client0.ObjectGetRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockClient)(nil).GetObject), varargs...)
}

// HashObjectPayloadRanges mocks base method.
func (m *MockClient) HashObjectPayloadRanges(arg0 context.Context, arg1 *client0.RangeChecksumParams, arg2 ...client0.CallOption) (*client0.ObjectRangeHashRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HashObjectPayloadRanges", varargs...)
	ret0, _ := ret[0].(*client0.ObjectRangeHashRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashObjectPayloadRanges indicates an expected call of HashObjectPayloadRanges.
func (mr *MockClientMockRecorder) HashObjectPayloadRanges(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashObjectPayloadRanges", reflect.TypeOf((*MockClient)(nil).HashObjectPayloadRanges), varargs...)
}

// HeadObject mocks base method.
func (m *MockClient) HeadObject(arg0 context.Context, arg1 *client0.ObjectHeaderParams, arg2 ...client0.CallOption) (*client0.ObjectHeadRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadObject", varargs...)
	ret0, _ := ret[0].(*client0.ObjectHeadRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject.
func (mr *MockClientMockRecorder) HeadObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockClient)(nil).HeadObject), varargs...)
}

// ListContainers mocks base method.
func (m *MockClient) ListContainers(arg0 context.Context, arg1 client0.ContainerListPrm) (*client0.ContainerListRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "ListContainers", varargs...)
	ret0, _ := ret[0].(*client0.ContainerListRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockClientMockRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockClient)(nil).ListContainers), varargs...)
}

// NetworkInfo mocks base method.
func (m *MockClient) NetworkInfo(arg0 context.Context, arg1 client0.NetworkInfoPrm) (*client0.NetworkInfoRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "NetworkInfo", varargs...)
	ret0, _ := ret[0].(*client0.NetworkInfoRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkInfo indicates an expected call of NetworkInfo.
func (mr *MockClientMockRecorder) NetworkInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInfo", reflect.TypeOf((*MockClient)(nil).NetworkInfo), varargs...)
}

// ObjectPayloadRangeData mocks base method.
func (m *MockClient) ObjectPayloadRangeData(arg0 context.Context, arg1 *client0.RangeDataParams, arg2 ...client0.CallOption) (*client0.ObjectRangeRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectPayloadRangeData", varargs...)
	ret0, _ := ret[0].(*client0.ObjectRangeRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectPayloadRangeData indicates an expected call of ObjectPayloadRangeData.
func (mr *MockClientMockRecorder) ObjectPayloadRangeData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectPayloadRangeData", reflect.TypeOf((*MockClient)(nil).ObjectPayloadRangeData), varargs...)
}

// PutContainer mocks base method.
func (m *MockClient) PutContainer(arg0 context.Context, arg1 client0.ContainerPutPrm) (*client0.ContainerPutRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "PutContainer", varargs...)
	ret0, _ := ret[0].(*client0.ContainerPutRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutContainer indicates an expected call of PutContainer.
func (mr *MockClientMockRecorder) PutContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutContainer", reflect.TypeOf((*MockClient)(nil).PutContainer), varargs...)
}

// PutObject mocks base method.
func (m *MockClient) PutObject(arg0 context.Context, arg1 *client0.PutObjectParams, arg2 ...client0.CallOption) (*client0.ObjectPutRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*client0.ObjectPutRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockClientMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockClient)(nil).PutObject), varargs...)
}

// Raw mocks base method.
func (m *MockClient) Raw() *client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(*client.Client)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockClientMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockClient)(nil).Raw))
}

// SearchObjects mocks base method.
func (m *MockClient) SearchObjects(arg0 context.Context, arg1 *client0.SearchObjectParams, arg2 ...client0.CallOption) (*client0.ObjectSearchRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchObjects", varargs...)
	ret0, _ := ret[0].(*client0.ObjectSearchRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchObjects indicates an expected call of SearchObjects.
func (mr *MockClientMockRecorder) SearchObjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchObjects", reflect.TypeOf((*MockClient)(nil).SearchObjects), varargs...)
}

// SetEACL mocks base method.
func (m *MockClient) SetEACL(arg0 context.Context, arg1 client0.SetEACLPrm) (*client0.SetEACLRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	ret := m.ctrl.Call(m, "SetEACL", varargs...)
	ret0, _ := ret[0].(*client0.SetEACLRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetEACL indicates an expected call of SetEACL.
func (mr *MockClientMockRecorder) SetEACL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEACL", reflect.TypeOf((*MockClient)(nil).SetEACL), varargs...)
}
