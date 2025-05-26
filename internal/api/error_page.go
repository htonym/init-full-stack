package api

import "net/http"

type ErrorPage struct {
	BaseData
}

func (repo *HandlerRepo) ServerErrorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	template := NewTemplate()
	data := ErrorPage{BaseData: BaseData{ErrorOccurred: true}}
	template.Render(w, "server-error", data)
}

func (repo *HandlerRepo) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	template := NewTemplate()
	data := ErrorPage{BaseData: BaseData{ErrorOccurred: true}}
	template.Render(w, "not-found-error", data)
}
