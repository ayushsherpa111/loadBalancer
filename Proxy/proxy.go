package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type endPoint struct {
	address string

	revProxy *httputil.ReverseProxy
}

func (p *endPoint) GetAddress() string {
	return p.address
}

func (p *endPoint) IsAlive() bool { return true }

func (p *endPoint) Serve(w http.ResponseWriter, r *http.Request) {
	p.revProxy.ServeHTTP(w, r)
}

func NewEndPoint(addr string) *endPoint {
	URI, err := url.Parse(addr)
	if err != nil {
		return nil
	}
	return &endPoint{
		address:  addr,
		revProxy: httputil.NewSingleHostReverseProxy(URI),
	}
}
