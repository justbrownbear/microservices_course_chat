package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) DeleteChat(ctx context.Context, chatID uint64) error {
	// Инициализируем соединение
	transaction, err := grpcApiInstance.dbConnection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := service_provider.NewWithTransaction(&transaction)

	// Выполняем бизнес-логику
	err = deleteChatHandler(ctx, serviceProvider, chatID)
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

func deleteChatHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, chatID uint64) error {
	chatService := serviceProvider.GetChatService()

	err := chatService.DeleteChat(ctx, chatID)
	if err != nil {
		return err
	}

	return nil
}
