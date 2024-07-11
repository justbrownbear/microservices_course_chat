package chat_controller

import (
	"google.golang.org/grpc"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

type controller struct {
	chat_v1.UnimplementedChatV1Server
}

// InitChatController registers the chat service with the provided gRPC server.
// It ensures that the server can handle requests for the chat service defined by chat_v1.
func InitChatController(grpcServer *grpc.Server) {
	chat_v1.RegisterChatV1Server(grpcServer, &controller{})
}
