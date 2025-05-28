package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
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
	data.CSRFToken = nosurf.Token(r)

	var err error
	data.Widgets, err = repo.DB.WidgetList(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching widgets: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "widget-list", data)
}

type DetailWidgetData struct {
	BaseData
	Title  string
	Widget *models.Widget
}

func (repo *HandlerRepo) detailWidgetsPage(w http.ResponseWriter, r *http.Request) {
	var data DetailWidgetData
	data.Init(r.Context())
	data.Title = "Widgets"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		repo.NotFoundPage(w, r)
		return
	}

	data.Widget, err = repo.DB.WidgetDetails(r.Context(), id)
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching widget %d details: %v", id, err))
		repo.ServerErrorPage(w, r)
		return
	}

	fmt.Println("data.Widget", data.Widget)

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	template := NewTemplate()
	template.Render(w, "widget-detail", data)
}

func (repo *HandlerRepo) createWidget(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	widget := models.Widget{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	_, err := repo.DB.WidgetCreate(r.Context(), widget)
	if err != nil {
		slog.Error(fmt.Sprintf("creating widget details: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	repo.listWidgetsPage(w, r)
}
