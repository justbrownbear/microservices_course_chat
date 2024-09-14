package grpc_api

import (
	"context"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	pg "github.com/justbrownbear/microservices_course_chat/internal/transaction_manager"
)

// GrpcAPI - Интерфейс gRPC API
type GrpcAPI interface {
	// *** Чаты ***
	// Создание чата
	CreateChat(ctx context.Context, in *chat_service.CreateChatRequest) (uint64, error)
	// Удаление чата
	DeleteChat(ctx context.Context, chatID uint64) error
	// Отправка сообщения
	SendMessage(ctx context.Context, in *chat_service.SendMessageRequest) error
}

type grpcAPI struct {
	txManager pg.TxManager
}

// InitGrpcAPI инициализирует gRPC API
func InitGrpcAPI(txManager pg.TxManager) GrpcAPI {
	return &grpcAPI{
		txManager: txManager,
	}
}
