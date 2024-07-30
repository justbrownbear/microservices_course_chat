package user_controller

import (
	"context"
	"log"

	user_service "github.com/justbrownbear/microservices_course_chat/internal/service/user"
	"github.com/justbrownbear/microservices_course_chat/pkg/user_v1"
)

// ***************************************************************************************************
// ***************************************************************************************************
func (s *controller) CreateUser(ctx context.Context, req *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	log.Printf("Create request fired: %v", req.String())

	payload := &user_service.CreateUserRequest{
		Nickname: req.GetNickname(),
	}

	userID, err := s.grpcAPI.CreateUser(ctx, payload)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	result := &user_v1.CreateUserResponse{Id: userID}

	return result, nil
}
