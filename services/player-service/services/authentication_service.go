package services

import (
	"net/http"

	"github.com/Stardome-Team/Stardome-API/services/player-service/models"
	"github.com/Stardome-Team/Stardome-API/services/player-service/repositories"
	"github.com/Stardome-Team/Stardome-API/services/player-service/utils"
	"github.com/gin-gonic/gin"

	"github.com/Stardome-Team/Stardome-API/libraries/go/jwt"
)

const (
	keyID             = "id"
	keyUserName       = "userName"
	keyDisplayName    = "displayName"
	keyAvatarURL      = "avatarUrl"
	keyAvatarBlurHash = "avatarBlurHash"
)

// AuthenticationService :
type AuthenticationService interface {
	AuthenticatePlayer(c *gin.Context, a *models.PlayerAuthentication) (interface{}, *models.ErrorParsing)
}

// NewAuthenticationService :
func NewAuthenticationService(pr repositories.PlayerRepository) AuthenticationService {
	return &service{
		playerRepository: pr,
	}
}

// AuthenticatePlayer :
func (s *service) AuthenticatePlayer(c *gin.Context, pa *models.PlayerAuthentication) (interface{}, *models.ErrorParsing) {

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

	payload := getAuthenticationTokenPayload(player)

	token := jwt.GenerateToken(&c.Request.Host, payload)

	return token, nil
}

func getAuthenticationTokenPayload(player *models.Player) map[string]interface{} {
	return map[string]interface{}{
		keyID:             player.ID,
		keyUserName:       player.UserName,
		keyDisplayName:    player.DisplayName,
		keyAvatarURL:      player.AvatarURL,
		keyAvatarBlurHash: player.AvatarBlurHash,
	}
}
