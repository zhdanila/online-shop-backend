-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    id          SERIAL PRIMARY KEY,
    buyer_id    INT REFERENCES buyers (id) ON DELETE CASCADE,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
