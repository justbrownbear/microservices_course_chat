package grpc_api

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	user_service "github.com/justbrownbear/microservices_course_chat/internal/service/user"
)

// GrpcAPI - Интерфейс gRPC API
type GrpcAPI interface {
	// *** Пользователи ***
	// Создание пользователя
	CreateUser(ctx context.Context, in *user_service.CreateUserRequest) (uint64, error)
	// Удаление пользователя
	DeleteUser(ctx context.Context, userID uint64) error

	// *** Чаты ***
	// Создание чата
	CreateChat(ctx context.Context, in *chat_service.CreateChatRequest) (uint64, error)
	// Удаление чата
	DeleteChat(ctx context.Context, chatID uint64) error
	// Отправка сообщения
	SendMessage(ctx context.Context, in *chat_service.SendMessageRequest) error
}

type grpcAPI struct {
	dbConnection *pgxpool.Pool
}

// InitGrpcAPI инициализирует gRPC API
func InitGrpcAPI(dbPool *pgxpool.Pool) GrpcAPI {
	return &grpcAPI{
		dbConnection: dbPool,
	}
}
