package user_controller

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/justbrownbear/microservices_course_chat/pkg/user_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (s *controller) DeleteUser(ctx context.Context, req *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("Create request fired: %v", req.String())
	result := &emptypb.Empty{}

	userID := req.GetUserId()
	err := s.grpcAPI.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("%v", err)
		return result, err
	}

	return result, nil
}
