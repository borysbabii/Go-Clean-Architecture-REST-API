DROP TABLE IF EXISTS news CASCADE;

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

CREATE INDEX IF NOT EXISTS news_title_id_idx ON news (title);