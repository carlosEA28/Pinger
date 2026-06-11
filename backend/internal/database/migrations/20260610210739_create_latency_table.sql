-- +goose Up
CREATE TABLE IF NOT EXISTS latency (
    id UUID PRIMARY KEY,
    monitor_id UUID NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    response_time_ms DOUBLE PRECISION NOT NULL,
    status_code INTEGER NOT NULL,
    dns_lookup_ms DOUBLE PRECISION,
    tcp_connect_ms DOUBLE PRECISION,
    ttfb_ms DOUBLE PRECISION,
    is_up BOOLEAN NOT NULL,
    FOREIGN KEY (monitor_id) REFERENCES monitors(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS latency;
