package user_repository

import (
	"context"
)

// UserRepository - Интерфейс репозитория пользователей
type UserRepository interface {
	CreateUser(ctx context.Context, nickname string) (int64, error)
	DeleteUser(ctx context.Context, id int64) error
}
