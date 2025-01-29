package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IndraNurfa/music-catalog/internal/models/spotify"
	"github.com/IndraNurfa/music-catalog/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			expectedBody: spotify.SearchResponse{
				Limit:  10,
				Offset: 0,
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 9,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b2735149c948fde506624246a684", "https://i.scdn.co/image/ab67616d00001e025149c948fde506624246a684", "https://i.scdn.co/image/ab67616d000048515149c948fde506624246a684"},
						AlbumName:        "POST HUMAN: SURVIVAL HORROR",
						ArtistsName:      []string{"Bring Me The Horizon", "BABYMETAL"},
						Explicit:         true,
						ID:               "7CAbF0By0Fpnbiu6Xn5ZF7",
						Name:             "Kingslayer (feat. BABYMETAL)",
					},
					{
						AlbumType:        "album",
						AlbumTotalTracks: 11,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273413697269620e16f4466f543", "https://i.scdn.co/image/ab67616d00001e02413697269620e16f4466f543", "https://i.scdn.co/image/ab67616d00004851413697269620e16f4466f543"},
						AlbumName:        "That's The Spirit",
						ArtistsName:      []string{"Bring Me The Horizon"},
						Explicit:         false,
						ID:               "6o39Ln9118FKTMbM4BvcEy",
						Name:             "Drown",
					},
				},
				Total: 906,
			},
			wantErr: false,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "kingslayer", 10, 1).Return(&spotify.SearchResponse{
					Limit:  10,
					Offset: 0,
					Items: []spotify.SpotifyTrackObject{
						{
							AlbumType:        "album",
							AlbumTotalTracks: 9,
							AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b2735149c948fde506624246a684", "https://i.scdn.co/image/ab67616d00001e025149c948fde506624246a684", "https://i.scdn.co/image/ab67616d000048515149c948fde506624246a684"},
							AlbumName:        "POST HUMAN: SURVIVAL HORROR",
							ArtistsName:      []string{"Bring Me The Horizon", "BABYMETAL"},
							Explicit:         true,
							ID:               "7CAbF0By0Fpnbiu6Xn5ZF7",
							Name:             "Kingslayer (feat. BABYMETAL)",
						},
						{
							AlbumType:        "album",
							AlbumTotalTracks: 11,
							AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273413697269620e16f4466f543", "https://i.scdn.co/image/ab67616d00001e02413697269620e16f4466f543", "https://i.scdn.co/image/ab67616d00004851413697269620e16f4466f543"},
							AlbumName:        "That's The Spirit",
							ArtistsName:      []string{"Bring Me The Horizon"},
							Explicit:         false,
							ID:               "6o39Ln9118FKTMbM4BvcEy",
							Name:             "Drown",
						},
					},
					Total: 906,
				}, nil)
			},
		},
		{
			name:               "failed",
			expectedStatusCode: 400,
			expectedBody:       spotify.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "kingslayer", 10, 1).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/tracks/search?query=kingslayer&pageSize=10&pageIndex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
