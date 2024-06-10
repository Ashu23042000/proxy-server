package main

import (
	"github.com/Ashu23042000/logger/logger"
	"github.com/Ashu23042000/proxy-server/server"
)

func main() {
	log := logger.New(nil, "debug")
	proxyServer := server.New(log, ":8080")
	proxyServer.Start()
}
