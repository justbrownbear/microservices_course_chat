

-- name: CreateUser :one
INSERT INTO public.users (nickname)
	VALUES ($1)
	RETURNING id;

-- name: DeleteUser :exec
UPDATE public.users
	SET
		is_deleted = TRUE,
		delete_timestamp = NOW()
	WHERE id = $1;
