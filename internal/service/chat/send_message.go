package chat_service

import (
	"context"

	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
)

// SendMessageRequest - запрос на отправку сообщения
type SendMessageRequest struct {
	ChatID  uint64
	UserID  uint64
	Message string
}

func (chatServiceInstance *chatService) SendMessage(ctx context.Context, in *SendMessageRequest) error {
	payload := chat_repository.SendMessageParams{
		ChatID:  int64(in.ChatID),
		UserID:  int64(in.UserID),
		Message: in.Message,
	}
	err := chatServiceInstance.repository.SendMessage(ctx, payload)

	return err
}
