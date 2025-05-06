package api

import (
	"net/http"
)

type HomePageData struct {
	BaseData
	Title string
}

func (repo *HandlerRepo) homePage(w http.ResponseWriter, r *http.Request) {
	var data HomePageData
	data.Init(r.Context())
	data.Title = "Home"

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "home", data)
}
