package app

import (
	"PinGo/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	router     *gin.Engine
	logService api.LogService
}

func NewServer(router *gin.Engine, logService api.LogService) *Server {
	return &Server{
		router:     router,
		logService: logService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
