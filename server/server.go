package server

import (
	"fmt"
	"net"

	"github.com/Ashu23042000/logger/logger"
	"github.com/Ashu23042000/proxy-server/cache"
)

type IServer interface {
	Start() error
	Stop()
}

type Server struct {
	log           logger.ILogger
	listenAddress string
	listner       net.Listener
	cache         cache.ICache
}

func New(log logger.ILogger, listenAddr string, cache cache.ICache) IServer {
	return &Server{
		log:           log,
		listenAddress: listenAddr,
		cache:         cache,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}
	s.listner = ln

	defer ln.Close()

	fmt.Printf("Server listening... on port %s\n", s.listenAddress)

	s.acceptConnection()

	return nil
}

func (s *Server) Stop() {}
