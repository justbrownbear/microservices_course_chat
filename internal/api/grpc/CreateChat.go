package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) CreateChat(ctx context.Context, in *chat_service.CreateChatRequest) (uint64, error) {
	// Инициализируем соединение
	transaction, err := grpcApiInstance.dbConnection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := service_provider.NewWithTransaction(&transaction)

	// Выполняем бизнес-логику
	result, err := createChatHandler(ctx, serviceProvider, in)
	if err != nil {
		// nolint:errcheck
		transaction.Rollback(ctx)
		return 0, err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func createChatHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, in *chat_service.CreateChatRequest) (uint64, error) {
	chatService := serviceProvider.GetChatService()

	chatID, err := chatService.CreateChat(ctx, in)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
