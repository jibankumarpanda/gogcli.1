package cmd

import (
	"context"

	"google.golang.org/api/chat/v1"

	"github.com/jibankumarpanda/gogcli.1/internal/googleapi"
)

var newChatService func(ctx context.Context, email string) (*chat.Service, error) = googleapi.NewChat
