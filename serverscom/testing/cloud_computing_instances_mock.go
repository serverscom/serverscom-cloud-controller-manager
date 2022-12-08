package serverscom_testing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	serverscom "github.com/serverscom/serverscom-go-client/pkg"
)

// MockCloudComputingInstancesService is a mock of CloudComputingInstancesService interface.
type MockCloudComputingInstancesService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudComputingInstancesServiceMockRecorder
}

// MockCloudComputingInstancesServiceMockRecorder is the mock recorder for MockCloudComputingInstancesService.
type MockCloudComputingInstancesServiceMockRecorder struct {
	mock *MockCloudComputingInstancesService
}

// NewMockCloudComputingInstancesService creates a new mock instance.
func NewMockCloudComputingInstancesService(ctrl *gomock.Controller) *MockCloudComputingInstancesService {
	mock := &MockCloudComputingInstancesService{ctrl: ctrl}
	mock.recorder = &MockCloudComputingInstancesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudComputingInstancesService) EXPECT() *MockCloudComputingInstancesServiceMockRecorder {
	return m.recorder
}

// ApproveUpgrade mocks base method.
func (m *MockCloudComputingInstancesService) ApproveUpgrade(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveUpgrade", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApproveUpgrade indicates an expected call of ApproveUpgrade.
func (mr *MockCloudComputingInstancesServiceMockRecorder) ApproveUpgrade(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveUpgrade", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).ApproveUpgrade), ctx, id)
}

// Collection mocks base method.
func (m *MockCloudComputingInstancesService) Collection() serverscom.Collection[serverscom.CloudComputingInstance] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection")
	ret0, _ := ret[0].(serverscom.Collection[serverscom.CloudComputingInstance])
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Collection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Collection))
}

// Create mocks base method.
func (m *MockCloudComputingInstancesService) Create(ctx context.Context, input serverscom.CloudComputingInstanceCreateInput) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Create), ctx, input)
}

// CreatePTRRecord mocks base method.
func (m *MockCloudComputingInstancesService) CreatePTRRecord(ctx context.Context, cloudInstanceID string, input serverscom.PTRRecordCreateInput) (*serverscom.PTRRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePTRRecord", ctx, cloudInstanceID, input)
	ret0, _ := ret[0].(*serverscom.PTRRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePTRRecord indicates an expected call of CreatePTRRecord.
func (mr *MockCloudComputingInstancesServiceMockRecorder) CreatePTRRecord(ctx, cloudInstanceID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePTRRecord", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).CreatePTRRecord), ctx, cloudInstanceID, input)
}

// Delete mocks base method.
func (m *MockCloudComputingInstancesService) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Delete), ctx, id)
}

// DeletePTRRecord mocks base method.
func (m *MockCloudComputingInstancesService) DeletePTRRecord(ctx context.Context, cloudInstanceID, ptrRecordID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePTRRecord", ctx, cloudInstanceID, ptrRecordID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePTRRecord indicates an expected call of DeletePTRRecord.
func (mr *MockCloudComputingInstancesServiceMockRecorder) DeletePTRRecord(ctx, cloudInstanceID, ptrRecordID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePTRRecord", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).DeletePTRRecord), ctx, cloudInstanceID, ptrRecordID)
}

// Get mocks base method.
func (m *MockCloudComputingInstancesService) Get(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Get), ctx, id)
}

// PTRRecords mocks base method.
func (m *MockCloudComputingInstancesService) PTRRecords(id string) serverscom.Collection[serverscom.PTRRecord] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PTRRecords", id)
	ret0, _ := ret[0].(serverscom.Collection[serverscom.PTRRecord])
	return ret0
}

// PTRRecords indicates an expected call of PTRRecords.
func (mr *MockCloudComputingInstancesServiceMockRecorder) PTRRecords(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PTRRecords", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).PTRRecords), id)
}

// PowerOff mocks base method.
func (m *MockCloudComputingInstancesService) PowerOff(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOff indicates an expected call of PowerOff.
func (mr *MockCloudComputingInstancesServiceMockRecorder) PowerOff(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).PowerOff), ctx, id)
}

// PowerOn mocks base method.
func (m *MockCloudComputingInstancesService) PowerOn(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOn", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOn indicates an expected call of PowerOn.
func (mr *MockCloudComputingInstancesServiceMockRecorder) PowerOn(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOn", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).PowerOn), ctx, id)
}

// Reinstall mocks base method.
func (m *MockCloudComputingInstancesService) Reinstall(ctx context.Context, id string, input serverscom.CloudComputingInstanceReinstallInput) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reinstall", ctx, id, input)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reinstall indicates an expected call of Reinstall.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Reinstall(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reinstall", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Reinstall), ctx, id, input)
}

// Rescue mocks base method.
func (m *MockCloudComputingInstancesService) Rescue(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rescue", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rescue indicates an expected call of Rescue.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Rescue(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rescue", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Rescue), ctx, id)
}

// RevertUpgrade mocks base method.
func (m *MockCloudComputingInstancesService) RevertUpgrade(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevertUpgrade", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevertUpgrade indicates an expected call of RevertUpgrade.
func (mr *MockCloudComputingInstancesServiceMockRecorder) RevertUpgrade(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertUpgrade", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).RevertUpgrade), ctx, id)
}

// Unrescue mocks base method.
func (m *MockCloudComputingInstancesService) Unrescue(ctx context.Context, id string) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unrescue", ctx, id)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unrescue indicates an expected call of Unrescue.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Unrescue(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unrescue", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Unrescue), ctx, id)
}

// Update mocks base method.
func (m *MockCloudComputingInstancesService) Update(ctx context.Context, id string, input serverscom.CloudComputingInstanceUpdateInput) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, input)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Update(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Update), ctx, id, input)
}

// Upgrade mocks base method.
func (m *MockCloudComputingInstancesService) Upgrade(ctx context.Context, id string, input serverscom.CloudComputingInstanceUpgradeInput) (*serverscom.CloudComputingInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", ctx, id, input)
	ret0, _ := ret[0].(*serverscom.CloudComputingInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockCloudComputingInstancesServiceMockRecorder) Upgrade(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockCloudComputingInstancesService)(nil).Upgrade), ctx, id, input)
}
