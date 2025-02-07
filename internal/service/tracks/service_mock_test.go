// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock_test.go -package=tracks
//

// Package tracks is a generated GoMock package.
package tracks

import (
	context "context"
	reflect "reflect"

	trackactivities "github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	spotify "github.com/IndraNurfa/music-catalog/internal/repository/spotify"
	gomock "go.uber.org/mock/gomock"
)

// MockspotifyOutbound is a mock of spotifyOutbound interface.
type MockspotifyOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockspotifyOutboundMockRecorder
	isgomock struct{}
}

// MockspotifyOutboundMockRecorder is the mock recorder for MockspotifyOutbound.
type MockspotifyOutboundMockRecorder struct {
	mock *MockspotifyOutbound
}

// NewMockspotifyOutbound creates a new mock instance.
func NewMockspotifyOutbound(ctrl *gomock.Controller) *MockspotifyOutbound {
	mock := &MockspotifyOutbound{ctrl: ctrl}
	mock.recorder = &MockspotifyOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockspotifyOutbound) EXPECT() *MockspotifyOutboundMockRecorder {
	return m.recorder
}

// GetRecommendation mocks base method.
func (m *MockspotifyOutbound) GetRecommendation(ctx context.Context, limit int, trackID string) (*spotify.SpotifySearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendation", ctx, limit, trackID)
	ret0, _ := ret[0].(*spotify.SpotifySearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendation indicates an expected call of GetRecommendation.
func (mr *MockspotifyOutboundMockRecorder) GetRecommendation(ctx, limit, trackID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendation", reflect.TypeOf((*MockspotifyOutbound)(nil).GetRecommendation), ctx, limit, trackID)
}

// Search mocks base method.
func (m *MockspotifyOutbound) Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query, limit, offset)
	ret0, _ := ret[0].(*spotify.SpotifySearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockspotifyOutboundMockRecorder) Search(ctx, query, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockspotifyOutbound)(nil).Search), ctx, query, limit, offset)
}

// MocktrackactivitiesRepository is a mock of trackactivitiesRepository interface.
type MocktrackactivitiesRepository struct {
	ctrl     *gomock.Controller
	recorder *MocktrackactivitiesRepositoryMockRecorder
	isgomock struct{}
}

// MocktrackactivitiesRepositoryMockRecorder is the mock recorder for MocktrackactivitiesRepository.
type MocktrackactivitiesRepositoryMockRecorder struct {
	mock *MocktrackactivitiesRepository
}

// NewMocktrackactivitiesRepository creates a new mock instance.
func NewMocktrackactivitiesRepository(ctrl *gomock.Controller) *MocktrackactivitiesRepository {
	mock := &MocktrackactivitiesRepository{ctrl: ctrl}
	mock.recorder = &MocktrackactivitiesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktrackactivitiesRepository) EXPECT() *MocktrackactivitiesRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MocktrackactivitiesRepository) Create(ctx context.Context, model trackactivities.TrackActivity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MocktrackactivitiesRepositoryMockRecorder) Create(ctx, model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MocktrackactivitiesRepository)(nil).Create), ctx, model)
}

// Get mocks base method.
func (m *MocktrackactivitiesRepository) Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userID, spotifyID)
	ret0, _ := ret[0].(*trackactivities.TrackActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MocktrackactivitiesRepositoryMockRecorder) Get(ctx, userID, spotifyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MocktrackactivitiesRepository)(nil).Get), ctx, userID, spotifyID)
}

// GetBulkSpotifyIDs mocks base method.
func (m *MocktrackactivitiesRepository) GetBulkSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBulkSpotifyIDs", ctx, userID, spotifyIDs)
	ret0, _ := ret[0].(map[string]trackactivities.TrackActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBulkSpotifyIDs indicates an expected call of GetBulkSpotifyIDs.
func (mr *MocktrackactivitiesRepositoryMockRecorder) GetBulkSpotifyIDs(ctx, userID, spotifyIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBulkSpotifyIDs", reflect.TypeOf((*MocktrackactivitiesRepository)(nil).GetBulkSpotifyIDs), ctx, userID, spotifyIDs)
}

// Update mocks base method.
func (m *MocktrackactivitiesRepository) Update(ctx context.Context, model trackactivities.TrackActivity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MocktrackactivitiesRepositoryMockRecorder) Update(ctx, model any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MocktrackactivitiesRepository)(nil).Update), ctx, model)
}
