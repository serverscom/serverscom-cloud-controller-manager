// Code generated by MockGen. DO NOT EDIT.
// Source: ./vendor/github.com/serverscom/serverscom-go-client/pkg/hosts.go
//
// Generated by this command:
//
//	mockgen --destination ./serverscom/testing/hosts_mock.go --package=serverscom_testing --source ./vendor/github.com/serverscom/serverscom-go-client/pkg/hosts.go
//

// Package serverscom_testing is a generated GoMock package.
package serverscom_testing

import (
	context "context"
	reflect "reflect"

	serverscom "github.com/serverscom/serverscom-go-client/pkg"
	gomock "go.uber.org/mock/gomock"
)

// MockHostsService is a mock of HostsService interface.
type MockHostsService struct {
	ctrl     *gomock.Controller
	recorder *MockHostsServiceMockRecorder
}

// MockHostsServiceMockRecorder is the mock recorder for MockHostsService.
type MockHostsServiceMockRecorder struct {
	mock *MockHostsService
}

// NewMockHostsService creates a new mock instance.
func NewMockHostsService(ctrl *gomock.Controller) *MockHostsService {
	mock := &MockHostsService{ctrl: ctrl}
	mock.recorder = &MockHostsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostsService) EXPECT() *MockHostsServiceMockRecorder {
	return m.recorder
}

// AbortReleaseForDedicatedServer mocks base method.
func (m *MockHostsService) AbortReleaseForDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortReleaseForDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortReleaseForDedicatedServer indicates an expected call of AbortReleaseForDedicatedServer.
func (mr *MockHostsServiceMockRecorder) AbortReleaseForDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortReleaseForDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).AbortReleaseForDedicatedServer), ctx, id)
}

// Collection mocks base method.
func (m *MockHostsService) Collection() serverscom.Collection[serverscom.Host] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection")
	ret0, _ := ret[0].(serverscom.Collection[serverscom.Host])
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockHostsServiceMockRecorder) Collection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockHostsService)(nil).Collection))
}

// CreateDedicatedServers mocks base method.
func (m *MockHostsService) CreateDedicatedServers(ctx context.Context, input serverscom.DedicatedServerCreateInput) ([]serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDedicatedServers", ctx, input)
	ret0, _ := ret[0].([]serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDedicatedServers indicates an expected call of CreateDedicatedServers.
func (mr *MockHostsServiceMockRecorder) CreateDedicatedServers(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDedicatedServers", reflect.TypeOf((*MockHostsService)(nil).CreateDedicatedServers), ctx, input)
}

// CreatePTRRecordForDedicatedServer mocks base method.
func (m *MockHostsService) CreatePTRRecordForDedicatedServer(ctx context.Context, id string, input serverscom.PTRRecordCreateInput) (*serverscom.PTRRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePTRRecordForDedicatedServer", ctx, id, input)
	ret0, _ := ret[0].(*serverscom.PTRRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePTRRecordForDedicatedServer indicates an expected call of CreatePTRRecordForDedicatedServer.
func (mr *MockHostsServiceMockRecorder) CreatePTRRecordForDedicatedServer(ctx, id, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePTRRecordForDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).CreatePTRRecordForDedicatedServer), ctx, id, input)
}

// CreateSBMServers mocks base method.
func (m *MockHostsService) CreateSBMServers(ctx context.Context, input serverscom.SBMServerCreateInput) ([]serverscom.SBMServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSBMServers", ctx, input)
	ret0, _ := ret[0].([]serverscom.SBMServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSBMServers indicates an expected call of CreateSBMServers.
func (mr *MockHostsServiceMockRecorder) CreateSBMServers(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSBMServers", reflect.TypeOf((*MockHostsService)(nil).CreateSBMServers), ctx, input)
}

// DedicatedServerConnections mocks base method.
func (m *MockHostsService) DedicatedServerConnections(id string) serverscom.Collection[serverscom.HostConnection] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedServerConnections", id)
	ret0, _ := ret[0].(serverscom.Collection[serverscom.HostConnection])
	return ret0
}

// DedicatedServerConnections indicates an expected call of DedicatedServerConnections.
func (mr *MockHostsServiceMockRecorder) DedicatedServerConnections(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedServerConnections", reflect.TypeOf((*MockHostsService)(nil).DedicatedServerConnections), id)
}

// DedicatedServerDriveSlots mocks base method.
func (m *MockHostsService) DedicatedServerDriveSlots(id string) serverscom.Collection[serverscom.HostDriveSlot] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedServerDriveSlots", id)
	ret0, _ := ret[0].(serverscom.Collection[serverscom.HostDriveSlot])
	return ret0
}

// DedicatedServerDriveSlots indicates an expected call of DedicatedServerDriveSlots.
func (mr *MockHostsServiceMockRecorder) DedicatedServerDriveSlots(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedServerDriveSlots", reflect.TypeOf((*MockHostsService)(nil).DedicatedServerDriveSlots), id)
}

// DedicatedServerNetworks mocks base method.
func (m *MockHostsService) DedicatedServerNetworks(id string) serverscom.Collection[serverscom.Network] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedServerNetworks", id)
	ret0, _ := ret[0].(serverscom.Collection[serverscom.Network])
	return ret0
}

// DedicatedServerNetworks indicates an expected call of DedicatedServerNetworks.
func (mr *MockHostsServiceMockRecorder) DedicatedServerNetworks(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedServerNetworks", reflect.TypeOf((*MockHostsService)(nil).DedicatedServerNetworks), id)
}

// DedicatedServerPTRRecords mocks base method.
func (m *MockHostsService) DedicatedServerPTRRecords(id string) serverscom.Collection[serverscom.PTRRecord] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedServerPTRRecords", id)
	ret0, _ := ret[0].(serverscom.Collection[serverscom.PTRRecord])
	return ret0
}

// DedicatedServerPTRRecords indicates an expected call of DedicatedServerPTRRecords.
func (mr *MockHostsServiceMockRecorder) DedicatedServerPTRRecords(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedServerPTRRecords", reflect.TypeOf((*MockHostsService)(nil).DedicatedServerPTRRecords), id)
}

// DedicatedServerPowerFeeds mocks base method.
func (m *MockHostsService) DedicatedServerPowerFeeds(ctx context.Context, id string) ([]serverscom.HostPowerFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicatedServerPowerFeeds", ctx, id)
	ret0, _ := ret[0].([]serverscom.HostPowerFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DedicatedServerPowerFeeds indicates an expected call of DedicatedServerPowerFeeds.
func (mr *MockHostsServiceMockRecorder) DedicatedServerPowerFeeds(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicatedServerPowerFeeds", reflect.TypeOf((*MockHostsService)(nil).DedicatedServerPowerFeeds), ctx, id)
}

// DeletePTRRecordForDedicatedServer mocks base method.
func (m *MockHostsService) DeletePTRRecordForDedicatedServer(ctx context.Context, hostID, ptrRecordID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePTRRecordForDedicatedServer", ctx, hostID, ptrRecordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePTRRecordForDedicatedServer indicates an expected call of DeletePTRRecordForDedicatedServer.
func (mr *MockHostsServiceMockRecorder) DeletePTRRecordForDedicatedServer(ctx, hostID, ptrRecordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePTRRecordForDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).DeletePTRRecordForDedicatedServer), ctx, hostID, ptrRecordID)
}

// GetDedicatedServer mocks base method.
func (m *MockHostsService) GetDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDedicatedServer indicates an expected call of GetDedicatedServer.
func (mr *MockHostsServiceMockRecorder) GetDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).GetDedicatedServer), ctx, id)
}

// GetKubernetesBaremetalNode mocks base method.
func (m *MockHostsService) GetKubernetesBaremetalNode(ctx context.Context, id string) (*serverscom.KubernetesBaremetalNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubernetesBaremetalNode", ctx, id)
	ret0, _ := ret[0].(*serverscom.KubernetesBaremetalNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKubernetesBaremetalNode indicates an expected call of GetKubernetesBaremetalNode.
func (mr *MockHostsServiceMockRecorder) GetKubernetesBaremetalNode(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubernetesBaremetalNode", reflect.TypeOf((*MockHostsService)(nil).GetKubernetesBaremetalNode), ctx, id)
}

// GetSBMServer mocks base method.
func (m *MockHostsService) GetSBMServer(ctx context.Context, id string) (*serverscom.SBMServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSBMServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.SBMServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSBMServer indicates an expected call of GetSBMServer.
func (mr *MockHostsServiceMockRecorder) GetSBMServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSBMServer", reflect.TypeOf((*MockHostsService)(nil).GetSBMServer), ctx, id)
}

// PowerCycleDedicatedServer mocks base method.
func (m *MockHostsService) PowerCycleDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerCycleDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerCycleDedicatedServer indicates an expected call of PowerCycleDedicatedServer.
func (mr *MockHostsServiceMockRecorder) PowerCycleDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerCycleDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).PowerCycleDedicatedServer), ctx, id)
}

// PowerOffDedicatedServer mocks base method.
func (m *MockHostsService) PowerOffDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOffDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOffDedicatedServer indicates an expected call of PowerOffDedicatedServer.
func (mr *MockHostsServiceMockRecorder) PowerOffDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOffDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).PowerOffDedicatedServer), ctx, id)
}

// PowerOnDedicatedServer mocks base method.
func (m *MockHostsService) PowerOnDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOnDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOnDedicatedServer indicates an expected call of PowerOnDedicatedServer.
func (mr *MockHostsServiceMockRecorder) PowerOnDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOnDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).PowerOnDedicatedServer), ctx, id)
}

// ReinstallOperatingSystemForDedicatedServer mocks base method.
func (m *MockHostsService) ReinstallOperatingSystemForDedicatedServer(ctx context.Context, id string, input serverscom.OperatingSystemReinstallInput) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReinstallOperatingSystemForDedicatedServer", ctx, id, input)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReinstallOperatingSystemForDedicatedServer indicates an expected call of ReinstallOperatingSystemForDedicatedServer.
func (mr *MockHostsServiceMockRecorder) ReinstallOperatingSystemForDedicatedServer(ctx, id, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReinstallOperatingSystemForDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).ReinstallOperatingSystemForDedicatedServer), ctx, id, input)
}

// ReleaseSBMServer mocks base method.
func (m *MockHostsService) ReleaseSBMServer(ctx context.Context, id string) (*serverscom.SBMServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseSBMServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.SBMServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseSBMServer indicates an expected call of ReleaseSBMServer.
func (mr *MockHostsServiceMockRecorder) ReleaseSBMServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseSBMServer", reflect.TypeOf((*MockHostsService)(nil).ReleaseSBMServer), ctx, id)
}

// ScheduleReleaseForDedicatedServer mocks base method.
func (m *MockHostsService) ScheduleReleaseForDedicatedServer(ctx context.Context, id string) (*serverscom.DedicatedServer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleReleaseForDedicatedServer", ctx, id)
	ret0, _ := ret[0].(*serverscom.DedicatedServer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleReleaseForDedicatedServer indicates an expected call of ScheduleReleaseForDedicatedServer.
func (mr *MockHostsServiceMockRecorder) ScheduleReleaseForDedicatedServer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleReleaseForDedicatedServer", reflect.TypeOf((*MockHostsService)(nil).ScheduleReleaseForDedicatedServer), ctx, id)
}
