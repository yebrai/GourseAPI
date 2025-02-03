package server

import (
	"GourseAPI/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server is running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())

}
