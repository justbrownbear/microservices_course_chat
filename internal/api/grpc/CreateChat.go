package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func (grpcApiInstance *grpcAPI) CreateChat(ctx context.Context, in *chat_service.CreateChatRequest) (uint64, error) {
	//
	result, err := withTransaction( ctx, grpcApiInstance.dbPool,
		func ( ctx context.Context, serviceProvider service_provider.ServiceProvider ) (interface{}, error) {
			chatService := serviceProvider.GetChatService()

			chatID, err := createChat( ctx, chatService, in )

			if err != nil {
				return nil, err
			}

			return chatID, nil
		} )

	if err != nil {
		return 0, err
	}


	// // Инициализируем соединение
	// transaction, err := grpcApiInstance.dbPool.BeginTx(ctx, pgx.TxOptions{})
	// if err != nil {
	// 	return 0, err
	// }

	// // Инициализируем сервис-провайдер
	// serviceProvider := getServiceProvider(&transaction)

	// // Выполняем бизнес-логику
	// result, err := createChatHandler(ctx, serviceProvider, in)
	// if err != nil {
	// 	// nolint:errcheck
	// 	transaction.Rollback(ctx)
	// 	return 0, err
	// }

	// err = transaction.Commit(ctx)
	// if err != nil {
	// 	return 0, err
	// }

	return result.(uint64), nil
}

// Handler - функция, которая выполняется в транзакции
type Handler func(ctx context.Context, serviceProvider service_provider.ServiceProvider) (interface{}, error)

func withTransaction( ctx context.Context, dbPool *pgxpool.Pool, fn Handler ) (interface{}, error) {
	// Инициализируем соединение
	transaction, err := dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	// Инициализируем сервис-провайдер
	serviceProvider := getServiceProvider(&transaction)

	// Выполняем бизнес-логику
	result, err := fn( ctx, serviceProvider )
	if err != nil {
		// nolint:errcheck
		transaction.Rollback(ctx)
		return nil, err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func createChat(ctx context.Context, chatService chat_service.ChatService, in *chat_service.CreateChatRequest) (uint64, error) {
	chatID, err := chatService.CreateChat(ctx, in)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}

// func createChatHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, in *chat_service.CreateChatRequest) (uint64, error) {
// 	chatService := serviceProvider.GetChatService()

// 	chatID, err := chatService.CreateChat(ctx, in)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return chatID, nil
// }
