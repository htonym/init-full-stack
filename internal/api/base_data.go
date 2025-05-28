package api

import (
	"context"

	"github.com/thofftech/init-full-stack/internal/models"
)

type BaseData struct {
	User          *models.User
	Token         string
	AppEnv        string
	HideNavbar    bool
	Authorized    bool
	ErrorOccurred bool
	CSRFToken     string
}

func (b *BaseData) Init(ctx context.Context) {
	var ok bool
	b.User, ok = ctx.Value(userKey).(*models.User)
	if ok {
		b.Authorized = true
	}
}
