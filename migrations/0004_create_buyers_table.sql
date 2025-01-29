-- +goose Up
-- +goose StatementBegin
CREATE TABLE buyers
(
    id      SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    name    VARCHAR(100) NOT NULL,
    phone   VARCHAR(20)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS buyers;
-- +goose StatementEnd
