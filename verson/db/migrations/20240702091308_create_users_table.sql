-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    email text unique not null,
    password text not null,
    username text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null
);
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd
