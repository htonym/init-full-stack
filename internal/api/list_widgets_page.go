package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/thofftech/init-full-stack/internal/models"
)

type ListWidgetsPageData struct {
	BaseData
	Title   string
	Widgets []models.Widget
}

func (repo *HandlerRepo) listWidgetsPage(w http.ResponseWriter, r *http.Request) {
	var data ListWidgetsPageData
	data.Init(r.Context())
	data.Title = "Widgets"

	var err error
	data.Widgets, err = repo.DB.AllWidgets(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching widgets: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "list-widgets", data)
}
