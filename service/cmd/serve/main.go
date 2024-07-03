package main

import (
	"log/slog"

	"github.com/paas-ok/service/app"
	"github.com/paas-ok/service/environment"
	"github.com/paas-ok/service/prusaslicer"
	"github.com/paas-ok/service/slicers"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	env := environment.Get()

	app := app.New()
	if engine := prusaslicer.New(*env.Slicer); engine != nil {
		slicer := slicers.New(engine)
		app.WithSlicer(slicer)
	}

	app.Run(*env.Host, env.Port)
}
