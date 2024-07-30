package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) DeleteUser(ctx context.Context, userID uint64) error {
	// Инициализируем соединение
	transaction, err := grpcApiInstance.dbConnection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := service_provider.NewWithTransaction(&transaction)

	// Выполняем бизнес-логику
	err = deleteUserHandler(ctx, serviceProvider, userID)
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

func deleteUserHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, userID uint64) error {
	userService := serviceProvider.GetUserService()

	err := userService.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
