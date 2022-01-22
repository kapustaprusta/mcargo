package spotify

import (
	"errors"
	"os"
)

type Config struct {
	clientId     string
	clientSecret string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Parse() error {
	c.clientId = os.Getenv("SPOTIFY_CLIENT_ID")
	if len(c.clientId) == 0 {
		return errors.New("spotify client id is empty")
	}

	c.clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	if len(c.clientSecret) == 0 {
		return errors.New("spotify client secret is empty")
	}

	return nil
}
