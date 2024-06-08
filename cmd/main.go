package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)


func main(){
	fx.New(
		fx.Provide(
			zap.NewDevelopment,
		),
		fx.Invoke(),
	).Run()
}