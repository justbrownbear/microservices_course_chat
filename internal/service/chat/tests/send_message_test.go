package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	chat_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/chat"
	chat_repository_mock "github.com/justbrownbear/microservices_course_chat/internal/repository/chat/mocks"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
)

func TestSendMessage(test *testing.T) {
	test.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx context.Context
		in *chat_service.SendMessageRequest
	}

	mc := minimock.NewController(test)
	defer test.Cleanup(mc.Finish)

	// Делаем залипухи
	ctx			:= context.Background()
	chatID			:= gofakeit.Uint64()
	userID			:= gofakeit.Uint64()
	message			:= gofakeit.BeerName()
	request			:= &chat_service.SendMessageRequest{
		ChatID: chatID,
		UserID: userID,
		Message: message,
	}
	payload := chat_repository.SendMessageParams{
		ChatID: int64( chatID ),
		UserID: int64( userID ),
		Message: message,
	}
	serviceError := fmt.Errorf("service error")

	// Объявим тип для функции, которая будет возвращать моки сервисов
	type grpcAPIMockFunction func(mc *minimock.Controller) grpc_api.GrpcAPI

	// Создаем набор тестовых кейсов
	tests := []struct {
		name			string
		args			args
		err				error
		serviceMock	grpcAPIMockFunction
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				in: request,
			},
			err: nil,
			serviceMock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок ChatRepository
				chatRepositoryMock := chat_repository_mock.NewChatRepositoryMock(mc)
				chatRepositoryMock.SendMessageMock.Expect(ctx, payload).Return(nil)

				// Инициализируем ChatService моком ChatRepository
				return chat_service.New(chatRepositoryMock)
			},
		},
		{
			name: "chat_repository.SendMessage() fail case",
			args: args{
				ctx: ctx,
				in: request,
			},
			err: serviceError,
			serviceMock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок ChatRepository
				chatRepositoryMock := chat_repository_mock.NewChatRepositoryMock(mc)
				chatRepositoryMock.SendMessageMock.Expect(ctx, payload).Return(serviceError)

				// Инициализируем ChatService моком ChatRepository
				return chat_service.New(chatRepositoryMock)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		test.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.serviceMock(mc)

			err := chatServiceMock.SendMessage(tt.args.ctx, tt.args.in);
			require.Equal(t, tt.err, err)
		})
	}
}
