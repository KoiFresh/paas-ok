package app

import (
	"encoding/json"
	"net/http"

	"github.com/paas-ok/service/materials"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 100 // 100 MB

func (app *App) Slice(res http.ResponseWriter, req *http.Request) {
	if app.slicer == nil {
		http.Error(res, "This service is not configured to slice", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	files := req.MultipartForm.File["files"]
	results, err := app.slicer.Slice(
		files,
		materials.PLA(),
	)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(results)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(response)
}
