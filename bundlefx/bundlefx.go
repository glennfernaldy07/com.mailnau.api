package bundlefx

import (
	"com.mailnau.api/config"
	"com.mailnau.api/user"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	r *mux.Router,
	cfg config.Config,
) {
	var errs = make(chan error)
	c := make(chan os.Signal, 1)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				var (
					httpAddr = flag.String("http.addr", cfg.GetString(config.ServerAddress), "HTTP listen address")
				)
				flag.Parse()

				go func() {
					signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
					errs <- fmt.Errorf("%s", <-c)
				}()

				go func() {
					fmt.Printf("Listening to %s", *httpAddr)
					errs <- http.ListenAndServe(*httpAddr, r)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				<-c

				return nil
			},
		},
	)
}

func rootRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}

// CoreModules core modules provided to fx
var CoreModules = fx.Options(
	config.Modules,
	fx.Provide(rootRouter),
)

// EntityModules entity modules provided to fx
var EntityModules = fx.Options(
	user.Modules,
	fx.Invoke(registerHooks),
)
