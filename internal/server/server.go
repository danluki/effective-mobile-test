package server

import (
	"github.com/danluki/effective-mobile-test/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	addr   string
}

func NewServer(cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		router: router,
		addr:   cfg.HttpServerAddress,
	}
}

func (s *Server) Run() error {
	return s.router.Run(s.addr)
}
