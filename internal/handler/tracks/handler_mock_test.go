// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go
//
// Generated by this command:
//
//	mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
//

// Package tracks is a generated GoMock package.
package tracks

import (
	context "context"
	reflect "reflect"

	spotify "github.com/IndraNurfa/music-catalog/internal/models/spotify"
	trackactivities "github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	gomock "go.uber.org/mock/gomock"
)

// Mockservice is a mock of service interface.
type Mockservice struct {
	ctrl     *gomock.Controller
	recorder *MockserviceMockRecorder
	isgomock struct{}
}

// MockserviceMockRecorder is the mock recorder for Mockservice.
type MockserviceMockRecorder struct {
	mock *Mockservice
}

// NewMockservice creates a new mock instance.
func NewMockservice(ctrl *gomock.Controller) *Mockservice {
	mock := &Mockservice{ctrl: ctrl}
	mock.recorder = &MockserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockservice) EXPECT() *MockserviceMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *Mockservice) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query, pageSize, pageIndex, userID)
	ret0, _ := ret[0].(*spotify.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockserviceMockRecorder) Search(ctx, query, pageSize, pageIndex, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*Mockservice)(nil).Search), ctx, query, pageSize, pageIndex, userID)
}

// UpsertTrackActivities mocks base method.
func (m *Mockservice) UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTrackActivities", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertTrackActivities indicates an expected call of UpsertTrackActivities.
func (mr *MockserviceMockRecorder) UpsertTrackActivities(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTrackActivities", reflect.TypeOf((*Mockservice)(nil).UpsertTrackActivities), ctx, userID, request)
}
