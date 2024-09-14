package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	chat_service_mock "github.com/justbrownbear/microservices_course_chat/internal/service/chat/mocks"
	service_provider_mock "github.com/justbrownbear/microservices_course_chat/internal/service_provider/mocks"
	"github.com/justbrownbear/microservices_course_chat/internal/transaction_manager"
	transaction_manager_mock "github.com/justbrownbear/microservices_course_chat/internal/transaction_manager/mocks"
)

func TestSendMessage(test *testing.T) {
	test.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx context.Context
		in  *chat_service.SendMessageRequest
	}

	mc := minimock.NewController(test)

	// Делаем залипухи
	ctx := context.Background()
	chatID := gofakeit.Uint64()
	userID := gofakeit.Uint64()
	message := gofakeit.BeerName()
	request := &chat_service.SendMessageRequest{
		ChatID:  chatID,
		UserID:  userID,
		Message: message,
	}
	serviceError := fmt.Errorf("service error")

	// Создаем набор тестовых кейсов
	tests := []struct {
		name string
		args args
		err  error
		mock func(mc *minimock.Controller) grpc_api.GrpcAPI
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				in:  request,
			},
			err: nil,
			mock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок TxManager
				txManagerMock := transaction_manager_mock.NewTxManagerMock(mc)
				txManagerMock.WithTransactionMock.Set(
					func(ctx context.Context, handler transaction_manager.Handler) error {
						chatServiceMock := chat_service_mock.NewChatServiceMock(mc)
						chatServiceMock.SendMessageMock.Expect(ctx, request).Return(nil)

						serviceProviderMock := service_provider_mock.NewServiceProviderMock(mc)
						serviceProviderMock.GetChatServiceMock.Return(chatServiceMock)

						return handler(ctx, serviceProviderMock)
					},
				)

				// Инициализируем GrpcAPI моком TxManager и ChatService
				return grpc_api.InitGrpcAPI(txManagerMock)
			},
		},
		{
			name: "chatService.SendMessage() fail case",
			args: args{
				ctx: ctx,
				in:  request,
			},
			err: serviceError,
			mock: func(mc *minimock.Controller) grpc_api.GrpcAPI {
				// Делаем мок TxManager
				txManagerMock := transaction_manager_mock.NewTxManagerMock(mc)
				txManagerMock.WithTransactionMock.Set(
					func(ctx context.Context, handler transaction_manager.Handler) error {
						chatServiceMock := chat_service_mock.NewChatServiceMock(mc)
						chatServiceMock.SendMessageMock.Expect(ctx, request).Return(serviceError)

						serviceProviderMock := service_provider_mock.NewServiceProviderMock(mc)
						serviceProviderMock.GetChatServiceMock.Return(chatServiceMock)

						return handler(ctx, serviceProviderMock)
					},
				)

				// Инициализируем GrpcAPI моком TxManager и ChatService
				return grpc_api.InitGrpcAPI(txManagerMock)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		test.Run(tt.name, func(t *testing.T) {
			grpcAPIMock := tt.mock(mc)

			err := grpcAPIMock.SendMessage(tt.args.ctx, tt.args.in)
			require.Equal(t, tt.err, err)
		})
	}
}
