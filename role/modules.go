package role

import (
	"com.mailnau.api/role/delivery/http"
	"com.mailnau.api/role/repository/mysql"
	"com.mailnau.api/role/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(http.NewServerOption),
	fx.Provide(http.NewEndpoint),
	fx.Provide(service.NewService),
	fx.Provide(mysql.NewRepository),
	fx.Invoke(http.NewHandler),
)
