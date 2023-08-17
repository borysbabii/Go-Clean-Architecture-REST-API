DROP TABLE IF EXISTS categories CASCADE;

CREATE TABLE categories
(
    category_id UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    user_id     UUID         NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    parent_id   UUID         NOT NULL REFERENCES categories (category_id) ON DELETE CASCADE,
    name        VARCHAR(250) NOT NULL CHECK ( name <> '' ),
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP             DEFAULT NULL
);