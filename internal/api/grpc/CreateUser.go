package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"

	user_service "github.com/justbrownbear/microservices_course_chat/internal/service/user"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) CreateUser(ctx context.Context, in *user_service.CreateUserRequest) (uint64, error) {
	// Инициализируем соединение
	transaction, err := grpcApiInstance.dbConnection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := service_provider.NewWithTransaction(&transaction)

	// Выполняем бизнес-логику
	result, err := createUserHandler(ctx, serviceProvider, in)
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

func createUserHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, in *user_service.CreateUserRequest) (uint64, error) {
	userService := serviceProvider.GetUserService()

	userID, err := userService.CreateUser(ctx, in)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
