-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS um;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS um;
-- +goose StatementEnd
