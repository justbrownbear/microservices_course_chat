package tests

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock"

	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/internal/api/grpc/mocks"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)

func TestCreateChat(t *testing.T) {
	t.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx context.Context
		in *chat_service.CreateChatRequest
	}

	mc := minimock.NewController(t)
	defer t.Cleanup(mc.Finish)

	// Делаем залипухи
	ctx			:= context.Background()
	chatID		:= gofakeit.Uint64()
	userID		:= gofakeit.Uint64()
	chatName	:= gofakeit.BeerName()
	request := &chat_service.CreateChatRequest{
		UserID:   userID,
		ChatName: chatName,
	}
	response := chatID

	// Объявим тип для функции, которая будет возвращать моки сервисов
	type grpcAPIMockFunction func(mc *minimock.Controller) grpc_api.GrpcAPI

	// Создаем набор тестовых кейсов
	tests := []struct {
		name			string
		args			args
		want			uint64
		err				error
		serviceMocks	grpcAPIMockFunction
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
				in: request,
			},
			want: response,
			err: nil,
			// serviceMocks: func(mc *minimock.Controller) grpc_api.GrpcAPI {
			serviceMocks: func(mc *minimock.Controller) grpc_api.GrpcAPI {

// ******** ЭТО НЕПРАВИЛЬНО, Я ТУТ ДОЛЖЕН ЗАМОКАТЬ chatService.CreateChat()
// ДЛЯ ЭТОГО Я ДОЛЖЕН ПРОИНИЦИАЛИЗИРОВАТЬ InitGrpcAPI C chatService

// НУЖНО МЕНЯТЬ АРХИТЕКТУРУ ПРИЛОЖЕНИЯ КАК ТУТ https://school.olezhek28.courses/pl/teach/control/lesson/view?id=329953142&editMode=0
// РЕПОЗИТОРИЙ ДОЛЖЕН ВЫЗЫВАТЬ НЕ СО СТАНДАРТНЫМ КЛИЕНТОМ ПГ, А С ОБЁРНУТЫМ
ПОСмотреть файл internal/api/grpc/CreateChat copy.go
получается, что нам нужно мокать отдельные методы сервис-провайдера,
чтобы возвращались замоканные сервисы, а на вход InitGrpcAPI() нужно
передавать функцию инициализации сервис-провайдера

				mock := mocks.NewGrpcAPIMock(mc)
				mock.CreateChatMock.Expect(ctx, request).Return(response, nil)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			grpcAPIMock := tt.serviceMocks(mc)



			// chatServiceMock := tt.serviceMocks(minimock.NewController(t))
			// chatServiceMock := tt.serviceMocks(mc)
			// Получается, тут нужно замокать БД и сервис-провайдер
			// api := grpc_api.InitGrpcAPI(nil)

			// mockServiceProvider := mocks.NewServiceProviderMock(mockServiceProvider)
			// mockServiceProvider.On("GetUserService").Return(chatServiceMock)

			// chatID, err := api.CreateChat(tt.args.ctx, tt.args.in)
			// require.Equal(t, tt.err, err)
			// require.Equal(t, tt.want, chatID)
		})
	}
}


func createChatHandler(ctx context.Context, serviceProvider service_provider.ServiceProvider, in *chat_service.CreateChatRequest) (uint64, error) {
	chatService := serviceProvider.GetChatService()

	chatID, err := chatService.CreateChat(ctx, in)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
