package main

import (
	"net/http"

	"github.com/charliegreeny/checkout/internal/config"
	"github.com/charliegreeny/checkout/internal/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)


func main(){
	fx.New(
		fx.Provide(
			zap.NewDevelopment,
			config.NewDb),
		fx.Invoke(router),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}


func router() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/product", routes.Item())
	r.Mount("/cart", routes.Cart())

	http.ListenAndServe(":8080", r)
}