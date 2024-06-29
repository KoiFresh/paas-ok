package app

import (
	"encoding/json"
	"net/http"

	"github.com/paas-ok/service/database"
)

func (app *App) GetAllMaterials(res http.ResponseWriter, req *http.Request) {
	filter := database.MaterialFilter{}
	if colors, ok := req.URL.Query()["color"]; ok {
		filter.Color = colors
	}
	if names, ok := req.URL.Query()["name"]; ok {
		filter.Name = names
	}

	mats, err := app.store.FindAllMaterials(filter)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(mats)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(payload)
}
