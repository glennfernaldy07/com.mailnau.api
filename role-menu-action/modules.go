package role_menu_action

import (
	"com.mailnau.api/role-menu-action/delivery/http"
	"com.mailnau.api/role-menu-action/repository/mysql"
	"com.mailnau.api/role-menu-action/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(http.NewServerOption),
	fx.Provide(http.NewEndpoint),
	fx.Provide(service.NewService),
	fx.Provide(mysql.NewRepository),
	fx.Invoke(http.NewHandler),
)
