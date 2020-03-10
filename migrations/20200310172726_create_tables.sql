-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart (
        id SERIAL PRIMARY KEY
    );

CREATE TABLE IF NOT EXISTS cart_item (
    id SERIAL PRIMARY KEY,
    product TEXT NOT NULL,
    quantity INT NOT NULL,
    fk_cart_id INT REFERENCES cart(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_item;
DROP TABLE IF EXISTS cart;
-- +goose StatementEnd
