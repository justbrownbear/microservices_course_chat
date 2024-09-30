// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package chat_repository

import (
	"context"
)

const createChat = `-- name: CreateChat :one
INSERT INTO public.chats (admin_user_id, name, create_user_id)
	VALUES ($1, $2, $1)
	RETURNING id
`

type CreateChatParams struct {
	AdminUserID int64
	Name        string
}

func (q *Queries) CreateChat(ctx context.Context, arg CreateChatParams) (int64, error) {
	row := q.db.QueryRow(ctx, createChat, arg.AdminUserID, arg.Name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteChat = `-- name: DeleteChat :exec
UPDATE public.chats
	SET
		is_deleted = TRUE,
		delete_timestamp = NOW()
	WHERE
		id = $1
`

func (q *Queries) DeleteChat(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteChat, id)
	return err
}

const sendMessage = `-- name: SendMessage :exec
INSERT INTO public.messages (chat_id, user_id, message)
	VALUES ($1, $2, $3)
`

type SendMessageParams struct {
	ChatID  int64
	UserID  int64
	Message string
}

func (q *Queries) SendMessage(ctx context.Context, arg SendMessageParams) error {
	_, err := q.db.Exec(ctx, sendMessage, arg.ChatID, arg.UserID, arg.Message)
	return err
}
