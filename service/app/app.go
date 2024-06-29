package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/paas-ok/service/database"
	"github.com/paas-ok/service/slicers"
)

type App struct {
	router *mux.Router
	slicer *slicers.Slicer
	store  *database.Store
}

func New() *App {
	store, err := database.NewStore("user", "password", "localhost", 3306, "printdb")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	return &App{
		router: mux.NewRouter(),
		store:  store,
	}
}

func (app *App) WithSlicer(slicer *slicers.Slicer) *App {
	app.slicer = slicer
	return app
}

func (app *App) Run(host string, port int) {
	app.router.HandleFunc("/orders", app.Orders).Methods("POST")
	app.router.HandleFunc("/cura:slice", app.Slice).Methods("POST")
	app.router.HandleFunc("/materials", app.GetAllMaterials).Methods("GET")

	addr := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handlers.CORS()(app.router),
		WriteTimeout: 5 * time.Minute,
		ReadTimeout:  5 * time.Minute,
	}

	slog.Info(fmt.Sprintf("Listen on http://%s/", addr))

	err := server.ListenAndServe()
	slog.Error(err.Error())
}
