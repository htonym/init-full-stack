package api

import (
	"net/http"
)

type ProfilePageData struct {
	BaseData
	Title string
}

func (repo *HandlerRepo) myProfilePage(w http.ResponseWriter, r *http.Request) {
	var data ProfilePageData
	data.Init(r.Context())
	data.Title = "My Profile"

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "profile", data)
}
