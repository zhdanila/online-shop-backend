-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id) ON DELETE CASCADE,
    product_id INT REFERENCES products (id) ON DELETE CASCADE,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
