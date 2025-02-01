package tracks

import (
	"context"
	"strconv"

	"github.com/IndraNurfa/music-catalog/internal/models/trackactivities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error {
	activity, err := s.trackactivitiesRepo.Get(ctx, userID, request.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get record from database")
		return err
	}

	if err == gorm.ErrRecordNotFound || activity == nil {
		err = s.trackactivitiesRepo.Create(ctx, trackactivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: strconv.Itoa(int(userID)),
			UpdatedBy: strconv.Itoa(int(userID)),
		})
		if err != nil {
			log.Error().Err(err).Msg("error create to database")
			return err
		}
		return nil
	}
	activity.IsLiked = request.IsLiked
	err = s.trackactivitiesRepo.Update(ctx, *activity)
	if err != nil {
		log.Error().Err(err).Msg("error update record in database")
		return err
	}
	return nil
}
