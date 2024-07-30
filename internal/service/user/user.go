package user_service

import (
	"context"

	user_repository "github.com/justbrownbear/microservices_course_chat/internal/repository/user"
)

// UserService - Интерфейс сервиса пользователей
type UserService interface {
	// Создание пользователя
	CreateUser(ctx context.Context, in *CreateUserRequest) (uint64, error)
	// Удаление пользователя
	DeleteUser(ctx context.Context, userID uint64) error
}

type userService struct {
	repository user_repository.UserRepository
}

// New инициализирует сервис пользователей
func New(repository user_repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}
