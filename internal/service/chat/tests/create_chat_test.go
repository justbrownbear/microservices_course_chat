package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	chat_repository_mock "github.com/justbrownbear/microservices_course_chat/internal/repository/chat/mocks"
	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
)

func TestCreateChat(test *testing.T) {
	test.Parallel()

	// Создаем структуру входных параметров
	type args struct {
		ctx context.Context
		in  *chat_service.CreateChatRequest
	}

	mc := minimock.NewController(test)

	// Делаем залипухи
	ctx := context.Background()
	chatID := gofakeit.Int64()
	userID := gofakeit.Uint64()
	chatName := gofakeit.BeerName()
	request := &chat_service.CreateChatRequest{
		UserID:   userID,
		ChatName: chatName,
	}
	response := uint64(chatID)
	serviceError := fmt.Errorf("service error")

	// Создаем набор тестовых кейсов
	tests := []struct {
		name string
		args args
		want uint64
		err  error
		mock func(mc *minimock.Controller) chat_service.ChatService
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				in:  request,
			},
			want: response,
			err:  nil,
			mock: func(mc *minimock.Controller) chat_service.ChatService {
				// Делаем мок ChatRepository
				chatRepositoryMock := chat_repository_mock.NewChatRepositoryMock(mc)
				chatRepositoryMock.CreateChatMock.Return(chatID, nil)

				// Инициализируем ChatService моком ChatRepository
				return chat_service.New(chatRepositoryMock)
			},
		},
		{
			name: "chat_repository.CreateChat() fail case",
			args: args{
				ctx: ctx,
				in:  request,
			},
			want: 0,
			err:  serviceError,
			mock: func(mc *minimock.Controller) chat_service.ChatService {
				// Делаем мок ChatRepository
				chatRepositoryMock := chat_repository_mock.NewChatRepositoryMock(mc)
				chatRepositoryMock.CreateChatMock.Return(0, serviceError)

				// Инициализируем ChatService моком ChatRepository
				return chat_service.New(chatRepositoryMock)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		test.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.mock(mc)

			chatID, err := chatServiceMock.CreateChat(tt.args.ctx, tt.args.in)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, chatID)
		})
	}
}
