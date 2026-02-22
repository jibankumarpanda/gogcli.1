package authclient

import (
	"context"
	"fmt"
	"strings"

	"github.com/jibankumarpanda/gogcli/internal/config"
)

type contextkey struct{}

func withclient(ctx context.Context, client string) context.Context {
	client = strings.TrimSpace(client)
	if client == "" {
		return ctx
	}
	return context.WithValue(ctx, contextkey{}, client)
}

func ClientOverrideFromContext(ctx context.Context) string {
	
}