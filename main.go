package main

import (
	"com.mailnau.api/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.CoreModules,
		bundlefx.EntityModules,
	).Run()
}
