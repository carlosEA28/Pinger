-- +goose Up
CREATE TABLE IF NOT EXISTS latency (
    id SERIAL PRIMARY KEY,
    monitor_id INTEGER NOT NULL,
    latency_ms INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (monitor_id) REFERENCES monitors(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS latency;
