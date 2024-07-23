-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE chats
(
    id         serial primary key,
    created_at timestamp not null default now(),
    updated_at timestamp
);

CREATE TABLE message
(
    id         serial primary key,
    author_id  bigint    not null,
    content    text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    chat_id    bigint references chats (id) on delete cascade
);

CREATE TABLE chat_members
(
    id      serial primary key,
    user_id bigint not null,
    chat_id bigint references chats (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table chats, chat_members, message;
-- +goose StatementEnd
