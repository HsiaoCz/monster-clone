-- +goose Up
CREATE TABLE IF NOT EXISTS sessionsas;
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS sessionsas;
-- +goose StatementBegin
-- +goose StatementEnd