-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);