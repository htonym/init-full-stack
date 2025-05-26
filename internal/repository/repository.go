package repository

import (
	"context"

	"github.com/thofftech/init-full-stack/internal/models"
)

type DatabaseRepo interface {
	AllWidgets(ctx context.Context) ([]models.Widget, error)
	WidgetDetails(ctx context.Context) (models.Widget, error)
}
