package main

import (
	"flag"
	"log"
	"os"
	"regexp"

	proxy "github.com/ayushsherpa111/loadBalancer/Proxy"
	"github.com/ayushsherpa111/loadBalancer/ldBlncr"
)

func parsePort(port string) bool {
	if match, err := regexp.Match(":[0-9]", []byte(port)); err != nil || !match {
		return false
	}
	return true
}

func main() {
	var port string
	var servers []ldBlncr.Server = make([]ldBlncr.Server, 0)
	flag.StringVar(&port, "p", "", "port where the load balancer listens on")

	flag.Parse()

	if len(port) == 0 || !parsePort(port) {
		log.Fatal("Invalid port provided. Port should be in the format -p (:[0-9])")
		os.Exit(1)
	}

	for _, v := range os.Args[3:] {
		servers = append(servers, proxy.NewEndPoint(v))
	}
	redirect := ldBlncr.NewLoadBalancer(port, servers...)

	redirect.HandleRoutes("/")

	redirect.Serve()
}
