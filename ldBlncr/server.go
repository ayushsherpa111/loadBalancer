package ldBlncr

import "net/http"

type Server interface {
	GetAddress() string

	IsAlive() bool

	Serve(http.ResponseWriter, *http.Request)
}
