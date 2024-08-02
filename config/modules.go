package config

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewViper),
	fx.Provide(NewMysqlDB),
)
