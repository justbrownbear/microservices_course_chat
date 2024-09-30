

-- name: CreateChat :one
INSERT INTO public.chats (admin_user_id, name, create_user_id)
	VALUES ($1, $2, $1)
	RETURNING id;

-- name: DeleteChat :exec
UPDATE public.chats
	SET
		is_deleted = TRUE,
		delete_timestamp = NOW()
	WHERE
		id = $1;

-- name: SendMessage :exec
INSERT INTO public.messages (chat_id, user_id, message)
	VALUES ($1, $2, $3);
