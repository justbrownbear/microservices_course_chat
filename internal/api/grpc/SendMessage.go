package grpc_api

import (
	"context"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/internal/service_provider"
)


func (instance *grpcAPI) SendMessage(
	ctx context.Context,
	in *chat_service.SendMessageRequest,
) error {
	err := instance.txManager.WithTransaction( ctx,
		func ( ctx context.Context, serviceProvider service_provider.ServiceProvider ) error {
			// В этом месте нам пришел сервис-провайдер, который уже имеет connection внутри себя
			// Нам осталось только получить нужные сервисы, и...
			chatService := serviceProvider.GetChatService()

			// ...Передать их функции, которая на входе принимает только используемые сервисы и in
			err := sendMessageHandler(ctx, chatService, in)
			if err != nil {
				return err
			}

			return nil
		} )
	if err != nil {
		return err
	}

	return nil;
}


func sendMessageHandler(
	ctx context.Context,
	chatService chat_service.ChatService,
	in *chat_service.SendMessageRequest,
) error {
	err := chatService.SendMessage(ctx, in)
	if err != nil {
		return err
	}

	return nil
}
