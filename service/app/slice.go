package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paas-ok/service/materials"
	"github.com/paas-ok/service/multipart"
	"github.com/paas-ok/service/slicers"
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

	form := multipart.Form(*req.MultipartForm)

	files := form.GetFiles("files")
	if len(files) == 0 {
		http.Error(res, "No files uploaded", http.StatusBadRequest)
		return
	}

	options := map[slicers.Option]string{}

	if infill, err := form.GetIntOrDefault("infill", -1); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	} else if infill >= 0 {
		options[slicers.OptionFillDensity] = fmt.Sprintf("%d%%", infill)
	}

	quality, err := form.GetStringEnumOrFirst("quality", []string{"low", "medium", "high"})
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	material, err := form.GetStringEnumOrFirst("material", []string{"pla", "abs", "petg", "tpu"})
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	options[slicers.OptionProfile] = fmt.Sprintf("%s_%s.ini", material, quality)

	results, err := app.slicer.Slice(
		files,
		materials.PLA(),
		options,
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
