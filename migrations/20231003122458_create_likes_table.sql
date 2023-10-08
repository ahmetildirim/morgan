-- +goose Up
-- +goose StatementBegin
CREATE TABLE likes (
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL,
    owner_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    CONSTRAINT unique_like_owner_id UNIQUE (post_id, owner_id)
);
-- +goose StatementEnd
