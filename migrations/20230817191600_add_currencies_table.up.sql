DROP TABLE IF EXISTS currencies CASCADE;

CREATE TABLE currencies
(
    currency_id UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    name        VARCHAR(250) NOT NULL CHECK ( name <> '' ),
    code        VARCHAR(3)   NOT NULL CHECK ( code <> '' ),
    symbol      VARCHAR(4)   NOT NULL CHECK ( symbol <> '' ),
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);