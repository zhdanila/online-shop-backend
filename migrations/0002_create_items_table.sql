-- +goose Up
-- +goose StatementBegin
CREATE TABLE items
(
    id          SERIAL PRIMARY KEY,
    seller_id   INT REFERENCES sellers (id) ON DELETE CASCADE,
    name        VARCHAR(255)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
