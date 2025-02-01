package memberships

import (
	"errors"

	"github.com/IndraNurfa/music-catalog/internal/models/memberships"
	"github.com/IndraNurfa/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetails, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error getting user from database")
		return "", err
	}

	if userDetails == nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	accessToken, err := jwt.CreateToken(userDetails.ID, userDetails.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create token")
	}

	return accessToken, nil
}
