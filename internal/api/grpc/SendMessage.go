package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) SendMessage(ctx context.Context, in *chat_service.SendMessageRequest) error {
	// Инициализируем соединение
	transaction, err := grpcApiInstance.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := getServiceProvider(&transaction)

	// Выполняем бизнес-логику
	err = sendMessageHandler(ctx, serviceProvider, in)
	if err != nil {
		// nolint:errcheck
		transaction.Rollback(ctx)
		return err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func sendMessageHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, in *chat_service.SendMessageRequest) error {
	chatService := serviceProvider.GetChatService()

	err := chatService.SendMessage(ctx, in)
	if err != nil {
		return err
	}

	return nil
}
