package user

import (
	"com.mailnau.api/user/delivery/http"
	"com.mailnau.api/user/repository/cache"
	"com.mailnau.api/user/repository/db"
	"com.mailnau.api/user/service"
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
