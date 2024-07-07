package chat_controller

import (
	"google.golang.org/grpc"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)


type controller struct {
	chat_v1.UnimplementedChatV1Server
}


func InitChatController( grpcServer *grpc.Server ) {

	chat_v1.RegisterChatV1Server( grpcServer, &controller{} )
}
