DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS news CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS budgets CASCADE;
DROP TABLE IF EXISTS goals CASCADE;
DROP TABLE IF EXISTS currencies CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS accounts CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE users
(
    user_id      UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    role         VARCHAR(10)                 NOT NULL DEFAULT 'user',
    about        VARCHAR(1024)                        DEFAULT '',
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    address      VARCHAR(250),
    city         VARCHAR(30),
    country      VARCHAR(30),
    gender       VARCHAR(20)                 NOT NULL DEFAULT 'male',
    postcode     INTEGER,
    birthday     DATE                                 DEFAULT NULL,
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP WITH TIME ZONE             DEFAULT NULL,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE news
(
    news_id    UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    author_id  UUID                     NOT NULL REFERENCES users (user_id),
    title      VARCHAR(250)             NOT NULL CHECK ( title <> '' ),
    content    TEXT                     NOT NULL CHECK ( content <> '' ),
    image_url  VARCHAR(1024) check ( image_url <> '' ),
    category   VARCHAR(250),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE          DEFAULT NULL
);

CREATE TABLE comments
(
    comment_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    author_id  UUID                                               NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    news_id    UUID                                               NOT NULL REFERENCES news (news_id) ON DELETE CASCADE,
    message    VARCHAR(1024)                                      NOT NULL CHECK ( message <> '' ),
    likes      BIGINT                   DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS news_title_id_idx ON news (title);

CREATE TABLE currencies
(
    currency_id UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    name        VARCHAR(250) NOT NULL CHECK ( name <> '' ),
    code        VARCHAR(3)   NOT NULL CHECK ( code <> '' ),
    symbol      VARCHAR(3)   NOT NULL CHECK ( symbol <> '' ),
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

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

CREATE TABLE budgets
(
    budget_id   UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id     UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    category_id UUID           NOT NULL REFERENCES categories (category_id) ON DELETE CASCADE,
    amount      NUMERIC(10, 2) NOT NULL,
    created_at  TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE goals
(
    goal_id    UUID PRIMARY KEY        DEFAULT uuid_generate_v4(),
    user_id    UUID           NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    name       VARCHAR(250)   NOT NULL CHECK ( name <> '' ),
    amount     NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);

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