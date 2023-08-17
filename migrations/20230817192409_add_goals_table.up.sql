DROP TABLE IF EXISTS goals CASCADE;

CREATE TABLE goals
(
    goal_id    UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id    UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    name       VARCHAR(250)   NOT NULL CHECK ( name <> '' ),
    amount     NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);