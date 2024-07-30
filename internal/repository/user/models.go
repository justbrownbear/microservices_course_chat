// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package user_repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Чаты
type Chat struct {
	ID              int64
	Name            string
	IsDeleted       bool
	AdminUserID     int64
	CreateUserID    int64
	CreateTimestamp pgtype.Timestamp
	UpdateUserID    pgtype.Int8
	UpdateTimestamp pgtype.Timestamp
	DeleteUserID    pgtype.Int8
	DeleteTimestamp pgtype.Timestamp
}

// Сообщения
type Message struct {
	ID              int64
	ChatID          int64
	UserID          int64
	Timestamp       pgtype.Timestamp
	Message         string
	IsDeleted       pgtype.Timestamp
	DeleteUserID    pgtype.Int8
	DeleteTimestamp pgtype.Timestamp
}

// Пользователи
type User struct {
	ID              int64
	Nickname        string
	IsDeleted       pgtype.Bool
	CreateTimestamp pgtype.Timestamp
	UpdateTimestamp pgtype.Timestamp
	DeleteTimestamp pgtype.Timestamp
}
