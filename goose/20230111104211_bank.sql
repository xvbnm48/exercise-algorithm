-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS banks;
CREATE TABLE banks (
    bank_id serial NOT NULL,
    bank_name varchar(255) NOT NULL,
    bank_icon varchar(255) NOT NULL,
    status varchar(255) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
    CONSTRAINT banks_pkey PRIMARY KEY (bank_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banks;
-- +goose StatementEnd
