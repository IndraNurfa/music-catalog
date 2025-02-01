package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/IndraNurfa/music-catalog/internal/models/spotify"
	"github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/IndraNurfa/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpotifyOutbound := NewMockspotifyOutbound(mockCtrl)
	mockTrackActivityRepo := NewMocktrackactivitiesRepository(mockCtrl)

	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=kingslayer&type=track&market=ID&locale=en-GB,en-US;q%3D0.9,en;q%3D0.8,id;q%3D0.7"
	islikedTrue := true
	islikedFalse := false
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "kingslayer",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotify.SearchResponse{
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
						IsLiked:          &islikedTrue,
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
						IsLiked:          &islikedFalse,
					},
				},
				Total: 906,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(&spotifyRepo.SpotifySearchResponse{
					Tracks: spotifyRepo.SpotifyTrack{
						Href:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=kingslayer&type=track&market=ID&locale=en-GB,en-US;q%3D0.9,en;q%3D0.8,id;q%3D0.7",
						Limit:  10,
						Next:   &next,
						Offset: 0,
						Total:  906,
						Items: []spotifyRepo.SpotifyTrackObject{
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 9,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
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
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:   "album",
									TotalTracks: 11,
									Images: []spotifyRepo.SpotifyAlbumImage{
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
								Artists: []spotifyRepo.SpotifyArtistObject{
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
				}, nil)
				mockTrackActivityRepo.EXPECT().GetBulkSpotifyIDs(gomock.Any(), uint(1), []string{"7CAbF0By0Fpnbiu6Xn5ZF7", "6o39Ln9118FKTMbM4BvcEy"}).
					Return(map[string]trackactivities.TrackActivity{
						"7CAbF0By0Fpnbiu6Xn5ZF7": {
							IsLiked: &islikedTrue,
						},
						"6o39Ln9118FKTMbM4BvcEy": {
							IsLiked: &islikedFalse,
						},
					}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:     "kingslayer",
				pageSize:  10,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound:     mockSpotifyOutbound,
				trackactivitiesRepo: mockTrackActivityRepo,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
