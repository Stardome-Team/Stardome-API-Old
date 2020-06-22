package utils

import (
	"testing"
	"time"

	"github.com/Blac-Panda/Stardome-API/models"
)

func TestGenerateToken(t *testing.T) {

	loadDevEnvironment()

	var playerID string = "playerID"
	var userName string = "userName"
	var emailAddress string = "test@test.com"
	var DisplayName string = "player display name"
	var avatarURL string = "localhost:8080/avatar.png"
	var avatarBlurHash string = "1821A8-A8ha91jd91"
	var timeNow time.Time = time.Now()

	tests := []struct {
		name   string
		player *models.Player
		host   string
	}{
		{
			name: "Token Generation 1",
			player: &models.Player{
				ID:             &playerID,
				UserName:       &userName,
				EmailAddress:   &emailAddress,
				DisplayName:    &DisplayName,
				AvatarURL:      &avatarURL,
				AvatarBlurHash: &avatarBlurHash,
				CreatedAt:      &timeNow,
				UpdatedAt:      nil,
				DeletedAt:      nil,
			},
			host: "STARDOME_TEST",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var token *models.Token = GenerateToken(&test.host, test.player)

			if token == nil {
				t.Errorf("utils.GenerateToken. \nError: Token object is %v+", token)
			}

			if len(*token.Token) == 0 {
				t.Errorf("utils.GenerateToken. \nError: Token length is %v+", len(*token.Token))
			}

			if token.Type != "Bearer" {
				t.Error("utils.GenerateToken, models.Token.Type != Bearer")
			}

			if timeNow.Add(time.Hour*1).UnixNano() > token.ExpiresOn {
				t.Errorf("utils.GenerateToken. \nError: Future Time %v+ | Generated Expirer Time = %v+ ", timeNow.Add(time.Hour*1).UnixNano(), token.ExpiresOn)
			}
		})
	}

}

func TestVerifyToken(t *testing.T) {

	loadDevEnvironment()

	tests := []struct {
		name     string
		token    string
		expected bool
	}{
		{
			name:     "Token Verification 1",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6Imp3dCJ9.eyJhdmF0YXJCbHVySGFzaCI6bnVsbCwiYXZhdGFyVXJsIjpudWxsLCJkaXNwbGF5TmFtZSI6bnVsbCwiZXhwIjoxNTkyMDY1Nzg1NTg1NzQ4ODAwLCJpYXQiOjE1OTIwNjIxODU1ODU3NDg4MDAsImlkIjoiYnF2ZzFvMzhiN29qanYxMWFiaDAiLCJpc3MiOiJsb2NhbGhvc3Q6MTAxMCIsIm5iZiI6MTU5MjA2MjE4ODU4NTc0ODgwMCwidXNlck5hbWUiOiJPbGEifQ.H1ixlqS46KoTwcOxmqfdZC0tKvoG9He-FueRp-tr1fU",
			expected: true,
		},
		{
			name:     "Token Verification 2",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := VerifyToken(test.token)

			if test.expected == true {
				if err != nil {
					t.Errorf("utils.VerifyToken. \nError: error = %v+", err)
				}
			} else {
				if err == nil {
					t.Errorf("utils.VerifyToken. \nError: verification failed")
				}
			}
		})
	}
}

func TestHasTokenExpired(t *testing.T) {

	loadDevEnvironment()

	var host string = "STARDOME_TEST"

	tests := []struct {
		name     string
		generate bool
		token    string
		expected bool
	}{
		{
			name:     "Token Expiration 1",
			generate: true,
			token:    "",
			expected: false,
		},
		{
			name:     "Token Expiration 2",
			generate: false,
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6Imp3dCJ9.eyJhdmF0YXJCbHVySGFzaCI6bnVsbCwiYXZhdGFyVXJsIjpudWxsLCJkaXNwbGF5TmFtZSI6bnVsbCwiZXhwIjoxNTkyMDY1Nzg1NTg1NzQ4ODAwLCJpYXQiOjE1OTIwNjIxODU1ODU3NDg4MDAsImlkIjoiYnF2ZzFvMzhiN29qanYxMWFiaDAiLCJpc3MiOiJsb2NhbGhvc3Q6MTAxMCIsIm5iZiI6MTU5MjA2MjE4ODU4NTc0ODgwMCwidXNlck5hbWUiOiJPbGEifQ.H1ixlqS46KoTwcOxmqfdZC0tKvoG9He-FueRp-tr1fU",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.generate {
				var token *models.Token = GenerateToken(&host, &models.Player{})

				if expired := HasTokenExpired(*token.Token); expired != test.expected {
					t.Errorf("utils.GenerateToken. \nError: expected = %v | result = %v ", test.expected, expired)
				}

			} else {

				if expired := HasTokenExpired(test.token); expired != test.expected {
					t.Errorf("utils.GenerateToken. \nError: expected = %v | result = %v ", test.expected, expired)
				}
			}
		})
	}
}
