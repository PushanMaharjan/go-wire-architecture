package internal

import (
	"go-wire-architecture/internal/user"
	"go.uber.org/fx"
)

var WebModule = fx.Module(
	"internals",
	fx.Options(
		user.Module,
	),
)
