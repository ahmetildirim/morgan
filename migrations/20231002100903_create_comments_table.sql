-- +goose Up
CREATE TABLE comments (
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL,
    owner_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);