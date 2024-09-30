package chat_controller

import (
	"google.golang.org/grpc"

	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

type controller struct {
	chat_v1.UnimplementedChatV1Server

	grpcAPI grpc_api.GrpcAPI
}

// InitChatController registers the chat service with the provided gRPC server.
// It ensures that the server can handle requests for the chat service defined by chat_v1.
func InitChatController(grpcServer *grpc.Server, grpcAPI grpc_api.GrpcAPI) {
	controllerInstance := &controller{
		grpcAPI: grpcAPI,
	}

	chat_v1.RegisterChatV1Server(grpcServer, controllerInstance)
}
