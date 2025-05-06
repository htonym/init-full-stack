package api

import (
	"context"

	"github.com/thofftech/init-full-stack/internal/auth"
)

type BaseData struct {
	User       *auth.User
	Token      string
	AppEnv     string
	HideNavbar bool
	Authorized bool
}

func (b *BaseData) Init(ctx context.Context) {
	var ok bool
	b.Token, ok = ctx.Value(tokenKey).(string)
	if ok {
		b.Authorized = true
	}
}
