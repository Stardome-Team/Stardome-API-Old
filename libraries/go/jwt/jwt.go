package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/jrpalma/jwt"
)

// Token :
type Token struct {
	Token     *string `json:"token"`
	ExpiresOn int64   `json:"expiresOn"`
	Type      string  `json:"type"`
}

const (
	keyTokenType = "Bearer"
)

// GenerateToken :
func GenerateToken(host *string, playload map[string]interface{}) *Token {

	token := jwt.NewJWT()

	issueAt, err := token.Claims.GetIssuedAt()
	notBefore := issueAt.Add(time.Second * 3)
	expiration := issueAt.Add(time.Hour * 1)

	token.Claims.SetNotBefore(notBefore)
	token.Claims.SetExpiration(expiration)
	token.Claims.SetIssuer(*host)

	for key, value := range playload {
		token.Claims.Set(key, value)
	}

	base64JWT, err := token.Sign(getTokenSecretKey())

	if err != nil {
		return nil
	}

	return &Token{
		Token:     &base64JWT,
		ExpiresOn: expiration.UnixNano(),
		Type:      keyTokenType,
	}
}

// VerifyToken :
func VerifyToken(token string) error {

	_jwt := jwt.NewJWT()

	err := _jwt.Verify(token, getTokenSecretKey())

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

func getTokenSecretKey() string {

	return viper.GetString("JWT_SECRET_KEY")
}
