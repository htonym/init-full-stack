package api

import (
	"fmt"
	"net/http"
)

type ListWidgetsPageData struct {
	BaseData
	Title string
}

func (repo *HandlerRepo) listWidgets(w http.ResponseWriter, r *http.Request) {
	var data ListWidgetsPageData
	data.Init(r.Context())
	data.Title = "Widgets"

	fmt.Printf("data: %+v", data)

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "list-widgets", data)
}
