package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Blac-Panda/Stardome-API/configurations"

	"github.com/Blac-Panda/Stardome-API/models"
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
func GenerateToken(host *string, player *models.Player) *models.Token {

	token := jwt.NewJWT()

	issueAt, err := token.Claims.GetIssuedAt()
	notBefore := issueAt.Add(time.Second * 3)
	expiration := issueAt.Add(time.Hour * 1)

	token.Claims.SetNotBefore(notBefore)
	token.Claims.SetExpiration(expiration)
	token.Claims.SetIssuer(*host)

	token.Claims.Set(keyID, player.ID)
	token.Claims.Set(keyUserName, player.UserName)
	token.Claims.Set(keyDisplayName, player.DisplayName)
	token.Claims.Set(keyAvatarURL, player.AvatarURL)
	token.Claims.Set(keyAvatarBlurHash, player.AvatarBlurHash)

	base64JWT, err := token.Sign(configurations.GetTokenSecretKey())

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

	err := _jwt.Verify(token, configurations.GetTokenSecretKey())

	if err != nil {
		return err
	}

	return nil
}

// HasTokenExpired :
func HasTokenExpired(token string) bool {

	tokenArr := strings.Split(token, ".")

	if len(tokenArr) != 3 {
		return false
	}

	claimsJSON, err := base64.RawURLEncoding.DecodeString(string(tokenArr[1]))

	if err != nil {
		return false
	}

	claimsObj := make(map[string]interface{})

	err = json.Unmarshal(claimsJSON, &claimsObj)

	if err != nil {
		fmt.Println(err)
	}

	expTime := int64(claimsObj["exp"].(float64))

	return expTime < time.Now().UnixNano()
}
