package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/paas-ok/service/slicers"
)

type App struct {
	router *mux.Router
	slicer *slicers.Slicer
}

func New() *App {
	return &App{
		router: mux.NewRouter(),
	}
}

func (app *App) WithSlicer(slicer *slicers.Slicer) *App {
	app.slicer = slicer
	return app
}

func (app *App) Run(host string, port int) {
	app.router.HandleFunc("/orders", app.Orders).Methods("POST")
	app.router.HandleFunc("/cura:slice", app.Slice).Methods("POST")

	addr := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{
		Addr:         addr,
		Handler:      app.router,
		WriteTimeout: 5 * time.Minute,
		ReadTimeout:  5 * time.Minute,
	}

	slog.Info(fmt.Sprintf("Listen on http://%s/", addr))

	err := server.ListenAndServe()
	slog.Error(err.Error())
}
