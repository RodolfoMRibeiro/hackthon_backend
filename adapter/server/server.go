package server

import (
	"fmt"
	"hackthon/adapter/env"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Run()
	GetServerEngine() *gin.Engine
}

type server struct {
	host, port   string
	serverEngine *gin.Engine
}

func CreateServer() IServer {
	server := NewServer(env.Server.HOST, env.Server.PORT)
	return server
}

func NewServer(host, port string) IServer {
	newServer := &server{
		host:         host,
		port:         port,
		serverEngine: gin.Default(),
	}
	return newServer
}

func (s *server) Run() {
	s.serverEngine.Run(fmt.Sprintf("%s:%s", s.host, s.port))
}

func (s *server) GetServerEngine() *gin.Engine {
	return s.serverEngine
}
