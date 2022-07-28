package server

import (
	"go-wire-architecture/pkg/server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userController *controllers.UserController, secondUserController *controllers.SecondUserController) *ServerHTTP {
	engine := gin.Default()

	engine.MaxMultipartMemory = 10485760

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	api := engine.Group("/api")

	// health check route
	api.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "API Up and Running"})
	})

	// user routes
	api.POST("/users", userController.Save)

	api.POST("/secondusers", secondUserController.Save)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":5000")
}
