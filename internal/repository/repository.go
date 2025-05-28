package repository

import (
	"context"

	"github.com/thofftech/init-full-stack/internal/models"
)

type DatabaseRepo interface {
	WidgetList(ctx context.Context) ([]models.Widget, error)
	WidgetDetails(ctx context.Context, id int) (*models.Widget, error)
	WidgetCreate(ctx context.Context, widget models.Widget) (models.Widget, error)
}
