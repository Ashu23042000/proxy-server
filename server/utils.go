package server

import (
	"bufio"
	"net"
	"strings"

	"github.com/Ashu23042000/proxy-server/model"
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

	requestParts := strings.Fields(request)

	if len(requestParts) != 3 {
		s.log.Error("Invalid request")
	}

	if strings.HasPrefix(request, "GET / ") {
		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"Hello, World!\r\n" +
			"\n"

		request := model.Request{
			Url:      requestParts[1],
			Response: response,
		}
		err := s.cache.InsertOne(request)
		if err != nil {
			s.log.Errorf("Error while inserting data into cache: %v", err)
		}

		res, err := s.cache.FindAll()
		if err != nil {
			s.log.Errorf("Error while getting data from cache: %v", err)
		}

		s.log.Debugf("request list: %v", res)

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
