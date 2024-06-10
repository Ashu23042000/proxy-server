package server

import (
	"fmt"
	"net"

	"github.com/Ashu23042000/logger/logger"
)

type IServer interface {
	Start() error
	Stop()
}

type Server struct {
	log           logger.ILogger
	listenAddress string
	listner       net.Listener
}

func New(log logger.ILogger, listenAddr string) IServer {
	return &Server{
		log:           log,
		listenAddress: listenAddr,
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
