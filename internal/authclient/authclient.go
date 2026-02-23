package authclient

import (
	"context"
	"strings"
)

type contextkey struct{}

func WithClient(ctx context.Context, client string) context.Context {
	client = strings.TrimSpace(client)
	if client == "" {
		return ctx
	}
	return context.WithValue(ctx, contextkey{}, client)
}

func ClientOverrideFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v := ctx.Value(contextkey{}); v != nil {
		return v.(string)
	}
	return ""
}
