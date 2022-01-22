package api

type Handler interface {
	Authenticate() error
}
