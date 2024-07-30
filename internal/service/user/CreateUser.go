package user_service

import (
	"context"
)

// CreateUserRequest - запрос на создание пользователя
type CreateUserRequest struct {
	Nickname string
}

// Создание пользователя
func (userServiceInstance *userService) CreateUser(ctx context.Context, in *CreateUserRequest) (uint64, error) {
	userID, err := userServiceInstance.repository.CreateUser(ctx, in.Nickname)
	if err != nil {
		return 0, err
	}

	return uint64(userID), err
}
