// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ZetoOfficial/neo4j-server/internal/repository (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -destination=mocks/repo_mock.go -package=mocks . Repository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/ZetoOfficial/neo4j-server/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddNodeAndRelationships mocks base method.
func (m *MockRepository) AddNodeAndRelationships(ctx context.Context, req models.AddNodeAndRelationshipsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeAndRelationships", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeAndRelationships indicates an expected call of AddNodeAndRelationships.
func (mr *MockRepositoryMockRecorder) AddNodeAndRelationships(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeAndRelationships", reflect.TypeOf((*MockRepository)(nil).AddNodeAndRelationships), ctx, req)
}

// DeleteNodeAndRelationships mocks base method.
func (m *MockRepository) DeleteNodeAndRelationships(ctx context.Context, nodeID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodeAndRelationships", ctx, nodeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodeAndRelationships indicates an expected call of DeleteNodeAndRelationships.
func (mr *MockRepositoryMockRecorder) DeleteNodeAndRelationships(ctx, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodeAndRelationships", reflect.TypeOf((*MockRepository)(nil).DeleteNodeAndRelationships), ctx, nodeID)
}

// GetAllNodes mocks base method.
func (m *MockRepository) GetAllNodes(ctx context.Context) ([]models.GetAllNodesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNodes", ctx)
	ret0, _ := ret[0].([]models.GetAllNodesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNodes indicates an expected call of GetAllNodes.
func (mr *MockRepositoryMockRecorder) GetAllNodes(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNodes", reflect.TypeOf((*MockRepository)(nil).GetAllNodes), ctx)
}

// GetAllRelationships mocks base method.
func (m *MockRepository) GetAllRelationships(ctx context.Context) ([]models.GetAllRelationshipsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRelationships", ctx)
	ret0, _ := ret[0].([]models.GetAllRelationshipsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRelationships indicates an expected call of GetAllRelationships.
func (mr *MockRepositoryMockRecorder) GetAllRelationships(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRelationships", reflect.TypeOf((*MockRepository)(nil).GetAllRelationships), ctx)
}

// GetNodeWithRelationships mocks base method.
func (m *MockRepository) GetNodeWithRelationships(ctx context.Context, nodeID int64) (models.NodeWithRelationships, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeWithRelationships", ctx, nodeID)
	ret0, _ := ret[0].(models.NodeWithRelationships)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeWithRelationships indicates an expected call of GetNodeWithRelationships.
func (mr *MockRepositoryMockRecorder) GetNodeWithRelationships(ctx, nodeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeWithRelationships", reflect.TypeOf((*MockRepository)(nil).GetNodeWithRelationships), ctx, nodeID)
}