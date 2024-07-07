package chat_controller

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"

)



func (s *controller) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {

	log.Printf("Delete request fired: %v", req.String())

	result := &emptypb.Empty{}

	return result, nil
}
