package user_controller

import (
	"google.golang.org/grpc"

	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/pkg/user_v1"
)

type controller struct {
	user_v1.UnimplementedUserV1Server

	grpcAPI grpc_api.GrpcAPI
}

// InitUserController инициализирует контроллер пользователя с предоставленным gRPC-сервером.
func InitUserController(grpcServer *grpc.Server, grpcAPI grpc_api.GrpcAPI) {
	controllerInstance := &controller{
		grpcAPI: grpcAPI,
	}

	user_v1.RegisterUserV1Server(grpcServer, controllerInstance)
}
