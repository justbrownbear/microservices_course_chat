package chat_service

import (
	"context"

	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
)

// ChatService - Интерфейс сервиса чатов
type ChatService interface {
	// Создание чата
	CreateChat(ctx context.Context, in *CreateChatRequest) (uint64, error)
	// Удаление чата
	DeleteChat(ctx context.Context, chatID uint64) error
	// Отправка сообщения
	SendMessage(ctx context.Context, in *SendMessageRequest) error
}

type chatService struct {
	repository chat_repository.ChatRepository
}

// New инициализирует сервис чатов
func New(repository chat_repository.ChatRepository) ChatService {
	return &chatService{
		repository: repository,
	}
}
