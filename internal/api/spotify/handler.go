package spotify

import (
	"net/http"
)

type Handler struct {
	config     *Config
	httpClient *http.Client
}

func NewHandler(config *Config, httpClient *http.Client) *Handler {
	return &Handler{
		config:     config,
		httpClient: httpClient,
	}
}
