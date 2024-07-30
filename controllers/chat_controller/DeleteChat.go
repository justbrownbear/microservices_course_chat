package chat_controller

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (chatControllerInstance *controller) DeleteChat(ctx context.Context, req *chat_v1.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("Delete request fired: %v", req.String())

	err := chatControllerInstance.grpcAPI.DeleteChat(ctx, req.ChatId)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &emptypb.Empty{}

	return result, nil
}
