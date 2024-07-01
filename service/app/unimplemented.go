package app

import "net/http"

func (app *App) Unimplemented(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
}
