package dbrepo

import (
	"context"
	"fmt"

	"github.com/thofftech/init-full-stack/internal/models"
)

func (m *postgresDBRepo) AllWidgets(ctx context.Context) ([]models.Widget, error) {
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

func (m *postgresDBRepo) WidgetDetails(ctx context.Context) (models.Widget, error) {
	return models.Widget{}, nil
}
