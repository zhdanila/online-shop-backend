-- +goose Up
-- +goose StatementBegin
CREATE TABLE buyers
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100) NOT NULL,
    phone VARCHAR(20)  NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS buyers;
-- +goose StatementEnd
