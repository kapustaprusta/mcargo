package spotify

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

var (
	authUrl     = "https://accounts.spotify.com/authorize"
	authScopes  = "user-library-read playlist-read-private"
	redirectUri = "http://localhost:10100/spotify_login_callback"

	allowedLetters = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ123456789!@#$%~"
)

type Loginer struct {
	config     *Config
	httpClient *http.Client
}

func NewLoginer(config *Config, httpClient *http.Client) *Loginer {
	return &Loginer{
		config:     config,
		httpClient: httpClient,
	}
}

func generateRandomString(stringLen int) string {
	rand.Seed(time.Now().UnixNano())
	randomBytes := make([]byte, stringLen)
	for byteIdx := 0; byteIdx < stringLen; byteIdx++ {
		letterIdx := rand.Intn(len(allowedLetters))
		randomBytes = append(randomBytes, allowedLetters[letterIdx])
	}

	return string(randomBytes)
}

func (l *Loginer) Url(ctx context.Context) (string, error) {
	url := authUrl + "?" +
		"show_dialog=true" +
		"response_type=code&" +
		"scope=" + authScopes + "&" +
		"redirect_uri=" + redirectUri + "&" +
		"client_id=" + l.config.clientId + "&" +
		"state=" + generateRandomString(16) + "&"

	return url, nil
}

func (l *Loginer) Login(ctx context.Context)
