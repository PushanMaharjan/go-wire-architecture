package bootstrap

import (
	"context"
	"go-wire-architecture/config"
	"go-wire-architecture/internal"
	"go-wire-architecture/pkg"
	"go.uber.org/fx"
)

func WebApplication() {
	app := fx.New(
		fx.Provide(
			config.New,
			pkg.GetLogger,
			pkg.NewHTTPServer,
			pkg.NewDatabase,
		),
		internal.WebModule,
		fx.Invoke(startWebServer),
	)

	app.Run()
}

func startWebServer(lifecycle fx.Lifecycle, server pkg.HTTPServer, config config.Config, logger pkg.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					server.Start(config, logger)
				}()
				return nil
			},
			OnStop: func(context context.Context) error {
				logger.Info("Stopping Application")
				return nil
			},
		},
	)
}
