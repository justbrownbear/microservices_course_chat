package chat_service

import (
	"context"

	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
)

// CreateChatRequest - запрос на создание чата
type CreateChatRequest struct {
	UserID   uint64
	ChatName string
}

func (chatServiceInstance *chatService) CreateChat(ctx context.Context, in *CreateChatRequest) (uint64, error) {
	payload := chat_repository.CreateChatParams{
		AdminUserID: int64(in.UserID),
		Name:        in.ChatName,
	}

	chatID, err := chatServiceInstance.repository.CreateChat(ctx, payload)
	if err != nil {
		return 0, err
	}

	return uint64(chatID), err
}
