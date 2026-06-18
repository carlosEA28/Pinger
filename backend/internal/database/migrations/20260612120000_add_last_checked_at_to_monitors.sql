-- +goose Up
ALTER TABLE monitors
    ADD COLUMN IF NOT EXISTS last_checked_at TIMESTAMP;

-- +goose Down
ALTER TABLE monitors
    DROP COLUMN IF EXISTS last_checked_at;
