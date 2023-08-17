DROP TABLE IF EXISTS accounts CASCADE;

CREATE TABLE accounts
(
    account_id      UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id         UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    currency_id     UUID           NOT NULL REFERENCES currencies (currency_id) ON DELETE CASCADE,
    name            VARCHAR(250)   NOT NULL CHECK ( name <> '' ),
    current_balance NUMERIC(10, 2) NOT NULL,
    created_at      TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);