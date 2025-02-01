package tracks

import (
	"context"

	"github.com/IndraNurfa/music-catalog/internal/models/spotify"
	"github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/IndraNurfa/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for index, item := range trackDetails.Tracks.Items {
		trackIDs[index] = item.ID
	}

	trackActivities, err := s.trackactivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil

}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for index, artists := range item.Artists {
			artistsName[index] = artists.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for index, image := range item.Album.Images {
			imageUrls[index] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			// album related fields
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,
			// artist related fields
			ArtistsName: artistsName,
			// track related fields
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
