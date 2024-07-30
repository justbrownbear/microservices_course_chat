package chat_controller

import (
	"context"
	"log"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (chatControllerInstance *controller) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (*chat_v1.CreateChatResponse, error) {
	log.Printf("Create request fired: %v", req.String())

	payload := &chat_service.CreateChatRequest{
		UserID:   req.GetUserId(),
		ChatName: req.GetChatName(),
	}

	chatID, err := chatControllerInstance.grpcAPI.CreateChat(ctx, payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &chat_v1.CreateChatResponse{
		Id: chatID,
	}

	return result, nil
}
