package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/IndraNurfa/music-catalog/internal/configs"
	"github.com/IndraNurfa/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)

	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=kingslayer&type=track&market=ID&locale=en-GB,en-US;q%3D0.9,en;q%3D0.8,id;q%3D0.7"

	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "kingslayer",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: SpotifyTrack{
					Href:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=kingslayer&type=track&market=ID&locale=en-GB,en-US;q%3D0.9,en;q%3D0.8,id;q%3D0.7",
					Limit:  10,
					Next:   &next,
					Offset: 0,
					Total:  906,
					Items: []SpotifyTrackObject{
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 9,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b2735149c948fde506624246a684",
									}, {
										URL: "https://i.scdn.co/image/ab67616d00001e025149c948fde506624246a684",
									}, {
										URL: "https://i.scdn.co/image/ab67616d000048515149c948fde506624246a684",
									},
								},
								Name: "POST HUMAN: SURVIVAL HORROR",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/1Ffb6ejR6Fe5IamqA5oRUF",
									Name: "Bring Me The Horizon",
								}, {
									Href: "https://api.spotify.com/v1/artists/630wzNP2OL7fl4Xl0GnMWq",
									Name: "BABYMETAL",
								},
							},
							Explicit: true,
							Href:     "https://api.spotify.com/v1/tracks/7CAbF0By0Fpnbiu6Xn5ZF7",
							ID:       "7CAbF0By0Fpnbiu6Xn5ZF7",
							Name:     "Kingslayer (feat. BABYMETAL)",
						},
						{
							Album: SpotifyAlbumObject{
								AlbumType:   "album",
								TotalTracks: 11,
								Images: []SpotifyAlbumImage{
									{
										URL: "https://i.scdn.co/image/ab67616d0000b273413697269620e16f4466f543",
									}, {
										URL: "https://i.scdn.co/image/ab67616d00001e02413697269620e16f4466f543",
									}, {
										URL: "https://i.scdn.co/image/ab67616d00004851413697269620e16f4466f543",
									},
								},
								Name: "That's The Spirit",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/1Ffb6ejR6Fe5IamqA5oRUF",
									Name: "Bring Me The Horizon",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/6o39Ln9118FKTMbM4BvcEy",
							ID:       "6o39Ln9118FKTMbM4BvcEy",
							Name:     "Drown",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := `https://api.spotify.com/v1/search`
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)

				req.Header.Set("Authorization", "Bearer accessToken")
				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       io.NopCloser(bytes.NewBufferString(`Internal Server Error`)),
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
