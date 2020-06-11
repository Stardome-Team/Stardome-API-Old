package utils

import (
	"time"

	"github.com/Blac-Panda/Stardome-API/configurations"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/gin-gonic/gin"
	"github.com/jrpalma/jwt"
)

const (
	keyTokenType = "Bearer"

	keyID             = "id"
	keyUserName       = "userName"
	keyDisplayName    = "displayName"
	keyAvatarURL      = "avatarUrl"
	keyAvatarBlurHash = "avatarBlurHash"
)

// GenerateToken :
func GenerateToken(c *gin.Context, player *models.Player) *models.Token {

	token := jwt.NewJWT()

	issueAt, err := token.Claims.GetIssuedAt()
	notBefore := issueAt.Add(time.Second * 3)
	expiration := issueAt.Add(time.Hour * 1)

	token.Claims.SetNotBefore(notBefore)
	token.Claims.SetExpiration(expiration)
	token.Claims.SetIssuer(c.Request.Host)

	token.Claims.Set(keyID, player.ID)
	token.Claims.Set(keyUserName, player.UserName)
	token.Claims.Set(keyDisplayName, player.DisplayName)
	token.Claims.Set(keyAvatarURL, player.AvatarURL)
	token.Claims.Set(keyAvatarBlurHash, player.AvatarBlurHash)

	base64JWT, err := token.Sign(configurations.JWTTokenSecretKey)

	if err != nil {
		return nil
	}

	return &models.Token{
		Token:     &base64JWT,
		ExpiresOn: expiration.UnixNano(),
		Type:      keyTokenType,
	}
}

// VerifyToken :
func VerifyToken(token string) error {

	_jwt := jwt.NewJWT()

	err := _jwt.Verify(token, configurations.JWTTokenSecretKey)

	if err != nil {
		return err
	}

	return nil
}
