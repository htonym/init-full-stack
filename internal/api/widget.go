package api

import (
	"encoding/json"
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
	data.CSRFToken = nosurf.Token(r)

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
	inputWidget := models.Widget{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	widget, err := repo.DB.WidgetCreate(r.Context(), inputWidget)
	if err != nil {
		slog.Error(fmt.Sprintf("creating widget details: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	http.Redirect(w, r, "/widgets/"+strconv.Itoa(widget.ID), http.StatusSeeOther)
}

func (repo *HandlerRepo) deleteWidget(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int `json:"id"`
	}

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Error decoding request body", "err", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate ID
	if req.ID == 0 {
		http.Error(w, "Missing or invalid id", http.StatusBadRequest)
		return
	}

	// Delete the widget
	err := repo.DB.WidgetDelete(r.Context(), req.ID)
	if err != nil {
		slog.Error(fmt.Sprintf("deleting widget: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	w.Header().Set("Location", "/widgets")
	w.WriteHeader(http.StatusNoContent)
}

func (repo *HandlerRepo) updateWidget(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Widget
	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		slog.Error("Error decoding request body", "err", err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate Input
	if reqBody.ID == 0 {
		http.Error(w, "Missing or invalid id", http.StatusBadRequest)
		return
	}

	_, err := repo.DB.WidgetUpdate(r.Context(), reqBody)
	if err != nil {
		slog.Error(fmt.Sprintf("deleting widget: %v", err))
		repo.ServerErrorPage(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}
