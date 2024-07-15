-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    user_id text unique not null,
    username text not null,
    email text unique not null,
    user_password text not null,
    synopsis text,
    avatar text not null,
    background_image text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    deleted_at timestamp with time zone 
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd