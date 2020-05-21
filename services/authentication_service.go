package services

import (
	"net/http"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
	"github.com/Blac-Panda/Stardome-API/utils"
	"github.com/gin-gonic/gin"
)

// AuthenticationService :
type AuthenticationService interface {
	AuthenticatePlayer(c *gin.Context, a *models.PlayerAuthentication) (*models.Token, *models.ErrorParsing)
}

// NewAuthenticationService :
func NewAuthenticationService(pr repositories.PlayerRepository) AuthenticationService {
	return &service{
		playerRepository: pr,
	}
}

// AuthenticatePlayer :
func (s *service) AuthenticatePlayer(c *gin.Context, pa *models.PlayerAuthentication) (*models.Token, *models.ErrorParsing) {

	player, err := s.playerRepository.GetPlayerByUserName(pa.UserName)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      utils.ErrorAuthenticationFailed,
			Type:       gin.ErrorTypePublic,
			Metadata:   utils.ReasonEntityNotFound,
			StatusCode: http.StatusNotFound,
		}
	}

	match, err := utils.CompareHashWithPassword(pa.Password, *player.PassHash)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      utils.ErrorInternalError,
			Type:       gin.ErrorTypePublic,
			Metadata:   utils.ReasonInternalServer,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !match {
		return nil, &models.ErrorParsing{
			Error:      utils.ErrorAuthenticationFailed,
			Type:       gin.ErrorTypePublic,
			Metadata:   utils.ReasonPasswordMismatch,
			StatusCode: http.StatusUnauthorized,
		}
	}

	token := utils.GenerateToken(c, player)

	return token, nil
}
