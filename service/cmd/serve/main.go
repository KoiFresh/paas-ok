package main

import (
	"flag"
	"log/slog"

	"github.com/paas-ok/service/app"
	"github.com/paas-ok/service/cura"
	"github.com/paas-ok/service/slicers"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	var fcura string
	flag.StringVar(&fcura, "cura", "CuraEngine", "The CureEngine binary path")

	var fport int
	flag.IntVar(&fport, "port", 8080, "The port to listen on")

	var fhost string
	flag.StringVar(&fhost, "host", "0.0.0.0", "The host to listen on")

	var fdir string
	flag.StringVar(&fdir, "config", "", "The directory used for configuration files")

	flag.Parse()

	app := app.New()
	if engine := cura.NewEngine(fcura, fdir); engine != nil {
		slicer := slicers.New(engine)
		app.WithSlicer(slicer)
	}

	app.Run(fhost, fport)
}
