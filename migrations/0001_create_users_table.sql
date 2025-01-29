-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(100) UNIQUE                                      NOT NULL,
    password   VARCHAR(255)                                             NOT NULL,
    role       VARCHAR(50) CHECK (role IN ('admin', 'seller', 'buyer')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
