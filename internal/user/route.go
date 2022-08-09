package user

import (
	"context"
	"go-wire-architecture/pkg"
	"go.uber.org/fx"
)

// Routes struct
type Routes struct {
	logger         pkg.Logger
	handler        pkg.HTTPServer
	userController *Handler
}

func NewRoutes(lc fx.Lifecycle, logger pkg.Logger, handler pkg.HTTPServer, userController *Handler) *Routes {
	routes := &Routes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Setting Up User Routes")
				routes.Setup()
				return nil
			},
		},
	)

	return routes
}

// Setup user routes
func (s *Routes) Setup() {
	s.logger.Info("Setting up routes")

	api := s.handler.Group("/api")

	api.GET("/user", s.userController.Save)

	api.POST("/user", s.userController.Save)
}
