package role

import (
	"com.mailnau.api/role/delivery/http"
	"com.mailnau.api/role/repository/cache"
	"com.mailnau.api/role/repository/db"
	"com.mailnau.api/role/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(http.NewServerOption),
	fx.Provide(http.NewEndpoint),
	fx.Provide(service.NewService),
	fx.Provide(db.NewRepository),
	fx.Provide(cache.NewRepository),
	fx.Invoke(http.NewHandler),
)
