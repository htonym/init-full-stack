package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/thofftech/init-full-stack/internal/models"
)

func (m *postgresDBRepo) WidgetList(ctx context.Context) ([]models.Widget, error) {
	sql := `SELECT id, name, description, created_at, updated_at FROM widgets`
	rows, err := m.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("query widgets: %w", err)
	}
	defer rows.Close()

	var widgets []models.Widget
	for rows.Next() {
		var w models.Widget
		if err := rows.Scan(&w.ID, &w.Name, &w.Description, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan widget: %w", err)
		}
		widgets = append(widgets, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate widgets: %w", err)
	}
	return widgets, nil
}

func (m *postgresDBRepo) WidgetDetails(ctx context.Context, pk int) (*models.Widget, error) {
	query := `
			SELECT
				w.id, w.name, w.description, w.created_at, w.updated_at,
				wc.id, wc.name, wc.widget_id, wc.complexity, wc.created_at, wc.updated_at
			FROM widgets w
			LEFT JOIN components wc ON wc.widget_id = w.id
			WHERE w.id = $1
			ORDER BY wc.id
		`

	rows, err := m.Pool.Query(ctx, query, pk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var widget *models.Widget
	components := []models.WidgetComponent{}

	for rows.Next() {
		var (
			widgetID      int
			widgetName    string
			widgetDesc    string
			widgetCreated time.Time
			widgetUpdated time.Time

			componentID         sql.NullInt64
			componentName       sql.NullString
			componentWidgetID   sql.NullInt64
			componentComplexity sql.NullInt64
			componentCreated    sql.NullString
			componentUpdated    sql.NullString
		)

		err := rows.Scan(
			&widgetID, &widgetName, &widgetDesc, &widgetCreated, &widgetUpdated,
			&componentID, &componentName, &componentWidgetID, &componentComplexity, &componentCreated, &componentUpdated,
		)
		if err != nil {
			return nil, err
		}

		// Initialize widget only once
		if widget == nil {
			widget = &models.Widget{
				ID:          widgetID,
				Name:        widgetName,
				Description: widgetDesc,
				CreatedAt:   widgetCreated,
				UpdatedAt:   widgetUpdated,
			}
		}

		// If componentID is valid, add component
		if componentID.Valid {
			component := models.WidgetComponent{
				ID:         int(componentID.Int64),
				Name:       componentName.String,
				WidgetID:   int(componentWidgetID.Int64),
				Complexity: int(componentComplexity.Int64),
				CreatedAt:  componentCreated.String,
				UpdatedAt:  componentUpdated.String,
			}
			components = append(components, component)
		}
	}

	if widget != nil {
		widget.Components = components
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if widget == nil {
		return nil, sql.ErrNoRows
	}

	return widget, nil
}

func (m *postgresDBRepo) WidgetCreate(ctx context.Context, widget models.Widget) (models.Widget, error) {
	var id int
	query := `INSERT INTO widgets (name, description) VALUES ($1, $2) RETURNING id`
	err := m.Pool.QueryRow(ctx, query, widget.Name, widget.Description).Scan(&id)
	if err != nil {
		return models.Widget{}, fmt.Errorf("insert new widget: %w", err)
	}
	widget.ID = id
	return widget, nil
}

func (m *postgresDBRepo) WidgetDelete(ctx context.Context, id int) error {
	query := `DELETE FROM widgets WHERE id = $1;`
	_, err := m.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("deleting widget: %w", err)
	}
	return nil
}
