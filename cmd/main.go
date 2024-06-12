package main

import (
	"github.com/charliegreeny/checkout/internal/api"
	"github.com/charliegreeny/checkout/internal/config"
	"github.com/charliegreeny/checkout/pkg/cart"
	"github.com/charliegreeny/checkout/pkg/item"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewDevelopment,
			config.NewDb,
			item.NewHandler,
			cart.NewHandler,
			cart.NewService,
		),
		fx.Invoke(api.StartRouter),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}