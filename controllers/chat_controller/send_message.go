package chat_controller

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	chat_service "github.com/justbrownbear/microservices_course_chat/internal/service/chat"
	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (chatControllerInstance *controller) SendMessage(ctx context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	payload := &chat_service.SendMessageRequest{
		ChatID:  req.ChatId,
		UserID:  req.UserId,
		Message: req.Message,
	}

	err := chatControllerInstance.grpcAPI.SendMessage(ctx, payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &emptypb.Empty{}

	return result, nil
}
