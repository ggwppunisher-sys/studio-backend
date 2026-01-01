-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS um.users (
    id UUID PRIMARY KEY,

    -- tg info
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    username VARCHAR(255),
    tg_id BIGINT,
    tg_chat_id BIGINT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS um.users;

-- +goose StatementEnd
