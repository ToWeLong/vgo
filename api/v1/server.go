package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/towelong/vgo/global"
	"github.com/towelong/vgo/middleware"
)

type Server struct {
	http *gin.Engine
}

func NewServer() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(
		middleware.Error,
		middleware.CORS,
		middleware.New(global.Logger).Log,
		middleware.Recovery(global.Logger),
	)
	srv := &Server{
		http: g,
	}
	srv.setupRouter()
	return srv.http
}

func (s *Server) setupRouter() {
	s.http.GET("/", s.ping)
}
