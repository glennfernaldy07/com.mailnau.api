package attendance

import (
	"com.mailnau.api/attendance/repository/cache"
	"com.mailnau.api/attendance/repository/db"
	"com.mailnau.api/attendance/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(service.NewService),
	fx.Provide(db.NewRepository),
	fx.Provide(cache.NewRepository),
)
