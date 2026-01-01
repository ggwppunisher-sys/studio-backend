-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS um.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    -- tg info
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(100),
    tg_id BIGINT,
    tg_chat_id BIGINT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS um.users;

-- +goose StatementEnd
