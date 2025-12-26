-- +goose Up
CREATE TABLE IF NOT EXISTS bot_statistics (
    id serial PRIMARY KEY,
    login text NOT NULL,
    request_name text NOT NULL,
    request_time timestamp NOT NULL DEFAULT now()
);