-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.chats
(
    id bigserial NOT NULL,
    name text NOT NULL,
    is_deleted boolean NOT NULL DEFAULT false,
    admin_user_id bigint NOT NULL,
    create_user_id bigint NOT NULL,
    create_timestamp timestamp without time zone NOT NULL DEFAULT NOW(),
    update_user_id bigint,
    update_timestamp timestamp without time zone,
    delete_user_id bigint,
    delete_timestamp timestamp without time zone,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.chats
    OWNER to postgres;

COMMENT ON TABLE public.chats
    IS 'Чаты';

CREATE TABLE public.messages
(
    id bigserial NOT NULL,
    chat_id bigint NOT NULL,
    user_id bigint NOT NULL,
    "timestamp" timestamp without time zone NOT NULL DEFAULT NOW(),
    message text NOT NULL,
    is_deleted timestamp without time zone,
    delete_user_id bigint,
    delete_timestamp timestamp without time zone,
    PRIMARY KEY (id),
    CONSTRAINT "messages_chat_id_FK" FOREIGN KEY (chat_id)
        REFERENCES public.chats (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
        NOT VALID
);

ALTER TABLE IF EXISTS public.messages
    OWNER to postgres;

COMMENT ON TABLE public.messages
    IS 'Сообщения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.messages;
DROP TABLE IF EXISTS public.chats;
-- +goose StatementEnd
