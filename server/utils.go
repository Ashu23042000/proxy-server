package server

import (
	"bufio"
	"net"
	"strings"
)

func (s *Server) acceptConnection() {
	for {
		conn, err := s.listner.Accept()
		if err != nil {
			s.log.Errorf("Error accepting connection: %v", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		s.log.Errorf("Error reading request: %v", err)
		return
	}

	s.log.Infof("Request recieved %s", string(request))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			s.log.Errorf("Error reading header line: %v", err)
			return
		}
		if line == "\r\n" || line == "\n" {
			break
		}
	}

	if strings.HasPrefix(request, "GET / ") {
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, World!\r\n" +
			"\n"
		conn.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 9\r\n" +
			"\r\n" +
			"Not Found"
		conn.Write([]byte(response))
	}

}
