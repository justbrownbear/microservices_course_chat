package chat_repository

import (
	"context"
)

// ChatRepository - Интерфейс репозитория чатов
type ChatRepository interface {
	CreateChat(ctx context.Context, arg CreateChatParams) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, arg SendMessageParams) error
}
