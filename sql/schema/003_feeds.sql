-- +goose Up

CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT Unique NOT NULL,
    user_id UUID references users(id) ON DELETE CASCADE
);


-- +goose Down

DROP TABLE feeds;