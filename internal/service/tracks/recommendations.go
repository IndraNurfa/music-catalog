package tracks

import (
	"context"
	"errors"

	"github.com/IndraNurfa/music-catalog/internal/models/spotify"
	"github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/IndraNurfa/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error) {
	if trackID == "" {
		return nil, errors.New("trackID cannot be empty")
	}

	trackDetails, err := s.spotifyOutbound.GetRecommendation(ctx, limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("error get recommendation request for spotify outbound in service")
		return nil, err
	}
	trackIDs := make([]string, len(trackDetails.Tracks))
	for index, item := range trackDetails.Tracks {
		trackIDs[index] = item.ID
	}

	trackActivities, err := s.trackactivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
		return nil, err
	}

	return modelToRecommendationResponse(trackDetails, trackActivities), nil
}

func modelToRecommendationResponse(data *spotifyRepo.SpotifyRecommendationResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.RecommendationResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks {
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

	return &spotify.RecommendationResponse{
		Items: items,
	}
}
