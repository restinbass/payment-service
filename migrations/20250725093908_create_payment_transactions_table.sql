-- +goose Up
-- +goose StatementBegin
CREATE TABLE payment_transactions (
    id uuid NOT NULL PRIMARY KEY,
    order_id uuid NOT NULL,
    user_id uuid NOT NULL,
    payment_method smallint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payment_transactions;
-- +goose StatementEnd
