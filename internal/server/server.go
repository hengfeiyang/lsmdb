package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hengfeiyang/lsmdb/internal/router"
)

type IServer interface {
	Run(addr string) error
}

type server struct{}

func New() IServer {
	return &server{}
}

func (s *server) Run(addr string) error {
	app := gin.Default()
	router.Route(app)
	return app.Run(addr)
}
