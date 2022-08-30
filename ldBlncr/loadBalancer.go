package ldBlncr

import (
	"log"
	"net/http"
	"os"
)

const (
	roundRobin uint8 = iota
	loadConnection
	responseTime
)

// Is responsible to acepting requests and forwarding it to available servers depending on how busy they are
type loadBalancer struct {
	port            string
	logger          *log.Logger
	servers         []Server
	requestQueue    map[string][]*http.Request
	mux             http.ServeMux
	roundRobinCount int
}

func (l *loadBalancer) HandleRoutes(route string) {
	l.mux.HandleFunc(route, l.loadBalance)
}

func (l *loadBalancer) loadBalance(w http.ResponseWriter, req *http.Request) {
	server := l.GetNextServer(roundRobin)
	server.Serve(w, req)
}

func (l *loadBalancer) GetNextServer(algo uint8) Server {
	var s Server
	switch algo {
	case roundRobin:
		s = l.algoRoundRobin()
	}
	return s
}

func (l *loadBalancer) Serve() {
	l.logger.Printf("Load balancer listening on port: [::]:%s\n", l.port)
	http.ListenAndServe(l.port, &l.mux)
}

func NewLoadBalancer(port string, servers ...Server) *loadBalancer {
	l := &loadBalancer{
		port:         port,
		servers:      servers,
		logger:       log.New(os.Stdout, " ", log.Lshortfile|log.Ltime),
		requestQueue: make(map[string][]*http.Request),
		mux:          *http.NewServeMux(),
	}

	for _, server := range servers {
		l.requestQueue[server.GetAddress()] = make([]*http.Request, 0)
	}

	return l
}
