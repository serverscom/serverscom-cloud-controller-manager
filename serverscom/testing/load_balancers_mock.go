package serverscom_testing

import (
	"github.com/golang/mock/gomock"
	pkg "github.com/serverscom/serverscom-go-client/pkg"
	"reflect"

	context "context"
)

// MockLoadBalancersService is a mock of LoadBalancersService interface.
type MockLoadBalancersService struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancersServiceMockRecorder
}

// MockLoadBalancersServiceMockRecorder is the mock recorder for MockLoadBalancersService.
type MockLoadBalancersServiceMockRecorder struct {
	mock *MockLoadBalancersService
}

// NewMockLoadBalancersService creates a new mock instance.
func NewMockLoadBalancersService(ctrl *gomock.Controller) *MockLoadBalancersService {
	mock := &MockLoadBalancersService{ctrl: ctrl}
	mock.recorder = &MockLoadBalancersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoadBalancersService) EXPECT() *MockLoadBalancersServiceMockRecorder {
	return m.recorder
}

// Collection mocks base method.
func (m *MockLoadBalancersService) Collection() pkg.Collection[pkg.LoadBalancer] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection")
	ret0, _ := ret[0].(pkg.Collection[pkg.LoadBalancer])
	return ret0
}

// Collection indicates an expected call of Collection.
func (mr *MockLoadBalancersServiceMockRecorder) Collection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockLoadBalancersService)(nil).Collection))
}

// CreateL4LoadBalancer mocks base method.
func (m *MockLoadBalancersService) CreateL4LoadBalancer(ctx context.Context, input pkg.L4LoadBalancerCreateInput) (*pkg.L4LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateL4LoadBalancer", ctx, input)
	ret0, _ := ret[0].(*pkg.L4LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateL4LoadBalancer indicates an expected call of CreateL4LoadBalancer.
func (mr *MockLoadBalancersServiceMockRecorder) CreateL4LoadBalancer(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateL4LoadBalancer", reflect.TypeOf((*MockLoadBalancersService)(nil).CreateL4LoadBalancer), ctx, input)
}

// DeleteL4LoadBalancer mocks base method.
func (m *MockLoadBalancersService) DeleteL4LoadBalancer(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteL4LoadBalancer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteL4LoadBalancer indicates an expected call of DeleteL4LoadBalancer.
func (mr *MockLoadBalancersServiceMockRecorder) DeleteL4LoadBalancer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteL4LoadBalancer", reflect.TypeOf((*MockLoadBalancersService)(nil).DeleteL4LoadBalancer), ctx, id)
}

// GetL4LoadBalancer mocks base method.
func (m *MockLoadBalancersService) GetL4LoadBalancer(ctx context.Context, id string) (*pkg.L4LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetL4LoadBalancer", ctx, id)
	ret0, _ := ret[0].(*pkg.L4LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetL4LoadBalancer indicates an expected call of GetL4LoadBalancer.
func (mr *MockLoadBalancersServiceMockRecorder) GetL4LoadBalancer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetL4LoadBalancer", reflect.TypeOf((*MockLoadBalancersService)(nil).GetL4LoadBalancer), ctx, id)
}

// UpdateL4LoadBalancer mocks base method.
func (m *MockLoadBalancersService) UpdateL4LoadBalancer(ctx context.Context, id string, input pkg.L4LoadBalancerUpdateInput) (*pkg.L4LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateL4LoadBalancer", ctx, id, input)
	ret0, _ := ret[0].(*pkg.L4LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateL4LoadBalancer indicates an expected call of UpdateL4LoadBalancer.
func (mr *MockLoadBalancersServiceMockRecorder) UpdateL4LoadBalancer(ctx, id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateL4LoadBalancer", reflect.TypeOf((*MockLoadBalancersService)(nil).UpdateL4LoadBalancer), ctx, id, input)
}
