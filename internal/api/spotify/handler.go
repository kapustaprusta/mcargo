package spotify

import (
	"context"
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

func (h *Handler) Authenticate(ctx context.Context) error {
	return nil
}
