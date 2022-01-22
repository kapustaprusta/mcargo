package spotify

import (
	"context"
	"net/http"
)

type Handler struct {
	httpClient *http.Client
}

func NewHandler(httpClient *http.Client) *Handler {
	return &Handler{
		httpClient: httpClient,
	}
}

func (h *Handler) Authenticate(ctx context.Context) error {
	return nil
}
