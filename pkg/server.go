package pkg

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	engine *gin.Engine
}

func NewHTTPServer() *HTTPServer {
	engine := gin.Default()

	engine.MaxMultipartMemory = 10485760

	engine.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
				AllowHeaders:     []string{"*"},
				AllowCredentials: true,
			},
		),
	)

	return &HTTPServer{engine: engine}
}

func (sh *HTTPServer) Start() {
	err := sh.engine.Run(":5000")
	if err != nil {
		return
	}
}
