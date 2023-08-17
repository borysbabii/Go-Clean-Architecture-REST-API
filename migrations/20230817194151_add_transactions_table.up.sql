DROP TABLE IF EXISTS transactions CASCADE;

CREATE TABLE transactions
(
    transaction_id UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id        UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    account_id     UUID           NOT NULL REFERENCES accounts (account_id) ON DELETE CASCADE,
    amount         NUMERIC(10, 2) NOT NULL,
    currency_id    UUID           NOT NULL REFERENCES currencies (currency_id) ON DELETE CASCADE,
    created_at     TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);