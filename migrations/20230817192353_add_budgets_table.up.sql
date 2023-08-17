DROP TABLE IF EXISTS budgets CASCADE;

CREATE TABLE budgets
(
    budget_id   UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id     UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    category_id UUID           NOT NULL REFERENCES categories (category_id) ON DELETE CASCADE,
    amount      NUMERIC(10, 2) NOT NULL,
    created_at  TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);
