package pkg

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

type HTTP struct {
    engine *gin.Engine
}

func NewServerHTTP() *HTTP {
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
    
    return &HTTP{engine: engine}
}

func (sh *HTTP) Start() {
    err := sh.engine.Run(":5000")
    if err != nil {
	  return
    }
}
